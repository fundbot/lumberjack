package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//
// Conf : read and handle configuration
//
type conf struct {
	fundBaseURL string
}

// Read environment variables
func (c *conf) readEnvVariables() {
	c.fundBaseURL = c.readEvnVar("FUND_BASE_URL")
	fmt.Println(c)
}

func (c *conf) readEvnVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Please set %s environment variable\n", key)
		os.Exit(1)
	}
	return value
}

//
// Server : Load up a tiny Server
//
type server struct{}

// Start a server to respond to pings
func (s *server) startServer() {
	fmt.Println("Listening at port 8080")
	http.HandleFunc("/", s.handler)
	http.ListenAndServe(":8080", nil)
}

// Reply to ping
func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong %s!", r.URL.Path[1:])
}

//
// Main: read config and start downloading files
//
func main() {
	var c conf
	c.readEnvVariables()

	// go getFiles(c.fundBaseURL)

	startDispatcher(3)
	for i := 0; i < 10; i++ {
		addToDownloadQueue("Job"+strconv.Itoa(i), time.Duration(i)*time.Second)
	}

	var s server
	s.startServer()
}

// Call a function every x seconds
func getFiles(url string) {
	for {
		<-time.After(1 * time.Second)

		start := time.Date(2006, time.April, 3, 0, 0, 0, 0, time.UTC)
		startTime := start.Format("02-Jan-2006")

		formattedURL := fmt.Sprintf(url, startTime)
		fmt.Println(formattedURL)

		body, err := downloadFile(formattedURL)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(body)
	}
}

// Download a file
func downloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//
// Tutorial on workers
//

// Queue of all downloads that we need to do
var downloadQueue = make(chan downloadRequest, 5000)

// Queue of workers to process downloads
var workerQueue chan chan downloadRequest

// Work Request
type downloadRequest struct {
	Name  string
	Delay time.Duration
}

// Download Worker
type worker struct {
	id              int
	dWorkerDownload chan downloadRequest
	dWorkerQueue    chan chan downloadRequest
	quit            chan bool
	taskCompleted   int
}

// Add work to the downloadQueue
func addToDownloadQueue(name string, delay time.Duration) {
	download := downloadRequest{Name: name, Delay: delay}
	fmt.Printf("Adding download %s to queue\n", name)
	downloadQueue <- download
}

func startDispatcher(numWorkers int) {
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
			case download := <-worker.dWorkerDownload:
				// Receive a work request
				fmt.Printf("worker %d: received work request, delaying for %f seconds\n", worker.id, download.Delay.Seconds())
				// time.Sleep(download.Delay)
				body, _ := downloadFile("http://example.qutheory.io/json")
				fmt.Println(body)
				worker.taskCompleted++
				fmt.Printf("worker %d: Hello, %s! Work is completed (%d)\n", worker.id, download.Name, worker.taskCompleted)
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
