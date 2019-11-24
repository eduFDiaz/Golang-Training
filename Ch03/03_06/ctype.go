// Writing a function that return Content-Type header
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// contentType will return the value of Content-Type header returned by making an
// HTTP GET request to url
func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() // Make sure we close the body

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" { // Return error if Content-Type header not found
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return ctype, nil
}

func getRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func main() {
	ctype, err := contentType("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}

	resp, err := getRequest("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(resp)
	}

	fmt.Println("Starting the application...")
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Println("Terminating the application...")
}
