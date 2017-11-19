package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type conf struct {
	fundBaseURL string
}

// Read environment variables
func (c *conf) readEnvVariables() {

	c.fundBaseURL = os.Getenv("FUND_BASE_URL")
	if c.fundBaseURL == "" {
		fmt.Println("Please set FUND_BASE_URL environment variable")
		os.Exit(1)
	}
	fmt.Println(c)

}

func main() {

	var configuration conf
	configuration.readEnvVariables()

	go getFiles()

	fmt.Println("Listening at port 8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// Reply to ping
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong %s!", r.URL.Path[1:])
}

// Call a function every x seconds
func getFiles() {
	for {
		<-time.After(1 * time.Second)

		url := "http://example.qutheory.io/json"

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
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
