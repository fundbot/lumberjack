package main

import (
	"github.com/fundbot/lumberjack/config"
	"github.com/fundbot/lumberjack/server"
	"github.com/fundbot/lumberjack/sync"
)

//
// Main: read config and start downloading files
//
func main() {
	config.SetVersion("0.1.0")
	config.Load()
	sync.StartProcessing()
	server.StartServer()
}
