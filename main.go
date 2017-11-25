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
	config.Load()
	sync.StartProcessing()
	server.StartServer()
}
