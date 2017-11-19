package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	go getFiles(c.fundBaseURL)

	var s server
	s.startServer()
}

// Call a function every x seconds
func getFiles(url string) {
	for {
		<-time.After(1 * time.Second)

		// url := "http://example.qutheory.io/json"

		body, err := download(url)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(body)
	}
}

// Download a file
func download(url string) (string, error) {
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
