package main

import (
	"fmt"
)

// metodos com structs

type Client struct {
	ID     int
	Name   string
	Active bool
}

func (c *Client) Disable() {

	c.Active = false
	fmt.Printf("O cliente %s foi desativado \n", c.Name)
}

func main() {
	client := Client{
		ID:     1,
		Name:   "tau",
		Active: true,
	}
	fmt.Println(client)
	client.Disable()

	fmt.Println(client)

}
