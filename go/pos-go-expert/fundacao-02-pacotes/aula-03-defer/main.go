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
	// o defer vai atrasar sua chamada e executar ela no final de todo processo
	defer response.Body.Close()
	dataReadable, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	println(string(dataReadable))

	println("-------")

	//
	// imprime na ordem correta
	println("linha 1")
	println("linha 2")
	println("linha 3")

	println("-------")
	//
	// imprime uma linha na ordem errada
	defer println("linha 1")
	println("linha 2")
	println("linha 3")

	// o defer vai executar somente nessa ultima parte do codigo
}
