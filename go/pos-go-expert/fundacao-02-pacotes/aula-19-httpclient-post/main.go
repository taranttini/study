package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	jsonVar := bytes.NewReader([]byte(`{"name":"taranttini"`))
	response, err := client.Post("https://google.com", "application/json", jsonVar)
	if err != nil {
		println("timeout ao acessar o site")
		return
	}
	defer response.Body.Close()
	body, err := io.CopyBuffer(os.Stdout, response.Body, nil)
	if err != nil {
		panic(err)
	}
	println(string(rune(body)))
}
