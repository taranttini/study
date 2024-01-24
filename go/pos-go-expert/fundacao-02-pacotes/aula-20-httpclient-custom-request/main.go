package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
		return
	}
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
