package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	go getFiles()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong %s!", r.URL.Path[1:])
}

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
