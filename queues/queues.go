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
	Key   string
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
func AddToDownloadQueue(date string, key string, delay time.Duration) {
	download := downloadRequest{Date: date, Key: key, Delay: delay}
	fmt.Printf("Adding download [%s] for processing\n", download.Key)
	downloadQueue <- download
}

// StartDispatcher : dispatch downloads to workers
func StartDispatcher(numWorkers int) {
	// Initialise the channel we are going to add worker's work into
	workerQueue = make(chan chan downloadRequest, numWorkers)

	// Create workers
	for i := 0; i < numWorkers; i++ {
		downloadWorker := newDownloadWorker(i+1, workerQueue)
		downloadWorker.Start()
	}
	fmt.Printf("[Dispatcher] Started %d workers\n", numWorkers)

	go func() {
		for {
			select {
			case download := <-downloadQueue:
				// fmt.Printf("Dispatcher : Received work request for %f seconds\n", download.Delay.Seconds())
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
			fmt.Printf("[Worker %d] Available for work again\n", worker.id)
			worker.dWorkerQueue <- worker.dWorkerDownload

			select {
			case downloadJob := <-worker.dWorkerDownload:
				// Receive a work request
				fmt.Printf("[Worker %d] Received work request, downloading %s\n", worker.id, downloadJob.Date)
				finalURL := fmt.Sprintf(config.BaseURL(), downloadJob.Date, downloadJob.Date)
				download.File(finalURL, downloadJob.Date, downloadJob.Key)
				// fmt.Println(body)
				fmt.Printf("[Worker %d] %s has been downloaded, off to sleep for %s\n", worker.id, downloadJob.Date, downloadJob.Delay.String())
				time.Sleep(downloadJob.Delay)
			case <-worker.quit:
				// Asked to stop working
				fmt.Printf("[Worker %d] Stopping\n", worker.id)
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
