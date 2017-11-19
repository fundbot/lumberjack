package main

import (
	"fmt"

	"github.com/fundbot/lumberjack/downloader"
)

func main() {

	body, err := downloader.Download("http://example.qutheory.io/json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)

}
