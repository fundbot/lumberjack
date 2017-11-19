package main

import (
	"io/ioutil"
	"net/http"
)

// Download : pass a URL and get a string of it's output
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
