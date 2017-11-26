package main

import (
	"github.com/sudhanshuraheja/lumberjack/config"
	"github.com/sudhanshuraheja/lumberjack/server"
	"github.com/sudhanshuraheja/lumberjack/sync"
)

//
// Main: read config and start downloading files
//
func main() {
	config.Load()
	sync.StartProcessing()
	server.StartServer()
}
