package main

import (
	"fmt"
)

// structs

type Client struct {
	ID     int
	Name   string
	Active bool
}

func main() {
	client := Client{
		ID:     1,
		Name:   "tau",
		Active: true,
	}
	fmt.Println(client)
	fmt.Printf("Nome: %s\n", client.Name)
	client.Active = false

	fmt.Printf("is active? %t\n", client.Active)
}
