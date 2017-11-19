package main

import (
	"fmt"
)

func main() {

	url := "http://example.qutheory.io/json"

	body, err := download(url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)

}
