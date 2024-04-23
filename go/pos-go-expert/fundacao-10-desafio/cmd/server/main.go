package main

import "github.com/taranttini/study/go/pos-go-expert/fundacao-10-desafio/configs"

func main() {
	println("ola mundo")
	config, _ := configs.LoadConfig(".")
	println(config.HttpTimeout)
	println(config.UrlBrasilApi)
	println(config.UrlViaCep)
	println(config.WebServerPort)
}
