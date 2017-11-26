package download

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// File : Download a file
func File(url string, date string, key string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	save(string(body), key)
	return string(body), nil
}

func save(data string, key string) (bool, error) {
	f, err := os.Create("./build/data/" + key)
	if err != nil {
		fmt.Println("[PANIC] Could not create file ./build/data" + key + " becuase of error " + err.Error())
		return false, err
	}
	defer f.Close()

	bytes, err := f.WriteString(data)
	if err != nil {
		fmt.Println("[PANIC] Could not write data to " + key + " because of error " + err.Error())
		return false, err
	}
	if bytes == 0 {
		fmt.Println("[PANIC] Could only write 0 bytes to " + key + " even though data has " + string(len(data)))
		return false, err
	}

	err = f.Sync()
	if err != nil {
		fmt.Println("[PANIC] f.Sync returned error " + err.Error())
		return false, err
	}

	return true, nil
}
