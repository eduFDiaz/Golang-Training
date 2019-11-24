package main

import (
	"fmt"
	"net/http"
)

func content(url string) (string, error) {
	resp, error := http.Get(url)
	content := resp.Header.Get("Content-Length")
	defer resp.Body.Close()
	return content, error
}

func main() {
	response, err := content("https://golang.org")
	if err != nil {
		fmt.Printf("Error: %s ", err)
	} else {
		fmt.Printf("response: %s ", response)
	}
}
