package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var url = "https://httpbin.org/get"

func main() {
	// send a request
	response, error := http.Get(url)
	if error != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", error)
	}

	// parse the body
	bytes, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Fprintf(os.Stderr, "parsing HTTP response body: %v\n", error)
		os.Exit(1)
	}
	response.Body.Close()

	// print the body out
	fmt.Printf("%s\n", bytes)
}
