package server

import (
	"fmt"
	"net/http"

	"github.com/fundbot/lumberjack/config"
)

// StartServer : to respond to pings, only running this because we might need it in the future
func StartServer() {
	connectString := fmt.Sprintf(":%d", config.Port())
	fmt.Printf("Application %s Listening at port %d\n", config.Name(), config.Port())
	http.HandleFunc("/", handler)
	http.ListenAndServe(connectString, nil)
}

// Reply to ping, sort of pointless
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong %s!", r.URL.Path[1:])
}
