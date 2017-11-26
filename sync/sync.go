package sync

import (
	"time"

	"github.com/fundbot/lumberjack/config"
	"github.com/fundbot/lumberjack/queues"
)

// StartProcessing : setup jobs and runners
func StartProcessing() {
	loadJobs()
	loadWorkers()
}

func loadJobs() {
	// Load Download Jobs
	// Start from 01-Apr-2006, [1143849600], and add a job for everyday till today
	end := time.Now().Unix()
	// end := int64(1144627200)
	start := int64(1143849600)
	day := int64(60 * 60 * 24)
	loop := start

	for loop < end {
		jobTime := time.Unix(loop, 0).Format("02-Jan-2006")
		key := time.Unix(loop, 0).Format("2006-01-02")
		queues.AddToDownloadQueue(jobTime, key, time.Duration(config.Delay())*time.Second)
		loop = loop + day
	}
}

func loadWorkers() {
	queues.StartDispatcher(1)
}
