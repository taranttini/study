package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: 1 * time.Microsecond}
	response, err := client.Get("https://google.com")
	if err != nil {
		println("timeout ao acessar o site")
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
