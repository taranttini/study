package main

import (
	"fmt"
)

// composicao de structs

type Client struct {
	ID     int
	Name   string
	Active bool
}

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

// essa nova estutura contempla o cliente e endereco (composicao)
type ClientType1 struct {
	Client
	Address
}

// essa nova estutura cria propriedade do tipo cliente e endereco
type ClientType2 struct {
	Client  Client
	Address Address
}

func main() {
	client1 := ClientType1{}
	client1.ID = 1
	client1.Name = "tau"
	client1.Active = true
	client1.Address.Street = "st one"
	client1.Address.Number = 1
	client1.Address.City = "Sao Paulo"
	client1.Address.State = "SP"

	fmt.Println(client1)

	client2 := ClientType2{
		Client: Client{
			ID:     1,
			Name:   "zezin",
			Active: true,
		},
		Address: Address{
			Street: "st zezin",
			Number: 10,
			City:   "Dream",
			State:  "DR",
		},
	}

	fmt.Println(client2)

}
