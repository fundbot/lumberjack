package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fundbot/lumberjack/config"
	"github.com/fundbot/lumberjack/download"
	"github.com/fundbot/lumberjack/queues"
	"github.com/fundbot/lumberjack/server"
)

//
// Main: read config and start downloading files
//
func main() {
	config.Load()
	// go getFiles(c.fundBaseURL)

	queues.StartDispatcher(3)
	for i := 0; i < 10; i++ {
		queues.AddToDownloadQueue("Job"+strconv.Itoa(i), time.Duration(i)*time.Second)
	}

	server.StartServer()
}

// Call a function every x seconds
func getFiles(url string) {
	for {
		<-time.After(1 * time.Second)

		start := time.Date(2006, time.April, 3, 0, 0, 0, 0, time.UTC)
		startTime := start.Format("02-Jan-2006")

		formattedURL := fmt.Sprintf(url, startTime)
		fmt.Println(formattedURL)

		body, err := download.File(formattedURL)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(body)
	}
}
