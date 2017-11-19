package main

import (
	"fmt"
)

func main() {

	body, err := download("http://example.qutheory.io/json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)

}
