package queues

import (
	"fmt"
	"time"

	"github.com/fundbot/lumberjack/config"
	"github.com/fundbot/lumberjack/download"
)

// Queue of all downloads that we need to do
var downloadQueue = make(chan downloadRequest, 5000)

// Queue of workers to process downloads
var workerQueue chan chan downloadRequest

// Work Request
type downloadRequest struct {
	Date  string
	Delay time.Duration
}

// Download Worker
type worker struct {
	id              int
	dWorkerDownload chan downloadRequest
	dWorkerQueue    chan chan downloadRequest
	quit            chan bool
}

// AddToDownloadQueue : Add work to the downloadQueue
func AddToDownloadQueue(date string, delay time.Duration) {
	download := downloadRequest{Date: date, Delay: delay}
	fmt.Printf("Adding download %s to queue\n", date)
	downloadQueue <- download
}

// StartDispatcher : dispatch downloads to workers
func StartDispatcher(numWorkers int) {
	// Initialise the channel we are going to add worker's work into
	workerQueue = make(chan chan downloadRequest, numWorkers)

	// Create workers
	for i := 0; i < numWorkers; i++ {
		fmt.Println("Dispathcher : Starting worker", i+1)
		downloadWorker := newDownloadWorker(i+1, workerQueue)
		downloadWorker.Start()
	}

	go func() {
		for {
			select {
			case download := <-downloadQueue:
				fmt.Printf("Dispatcher : Received work request for %f seconds\n", download.Delay.Seconds())
				go func() {
					downloader := <-workerQueue
					downloader <- download
				}()
			}
		}
	}()
}

// Create new download workers
func newDownloadWorker(id int, workerQueue chan chan downloadRequest) worker {
	downloadWorker := worker{
		id:              id,
		dWorkerDownload: make(chan downloadRequest),
		dWorkerQueue:    workerQueue,
		quit:            make(chan bool),
	}
	return downloadWorker
}

// Start a goroutine
func (worker *worker) Start() {
	go func() {
		for {
			// Add this worker to the workerQueue
			fmt.Printf("worker %d: adding worker back to workerQueue\n", worker.id)
			worker.dWorkerQueue <- worker.dWorkerDownload

			select {
			case downloadJob := <-worker.dWorkerDownload:
				// Receive a work request
				fmt.Printf("worker %d: received work request, delaying for %f seconds\n", worker.id, downloadJob.Delay.Seconds())
				finalURL := fmt.Sprintf(config.BaseURL(), downloadJob.Date, downloadJob.Date)
				body, _ := download.File(finalURL, downloadJob.Date)
				fmt.Println(body)
				fmt.Printf("worker %d: %s has been downloaded\n", worker.id, downloadJob.Date)
				time.Sleep(downloadJob.Delay)
			case <-worker.quit:
				// Asked to stop working
				fmt.Printf("worker %d stopping\n", worker.id)
				return
			}
		}
	}()
}

func (worker *worker) Stop() {
	go func() {
		worker.quit <- true
	}()
}
