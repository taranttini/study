package main

import (
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	dataReadable, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	println(string(dataReadable))

	response.Body.Close()
}
