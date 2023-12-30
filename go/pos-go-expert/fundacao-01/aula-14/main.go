package main

import (
	"fmt"
)

// interfaces

// interfaces em go, suportam somente metodos
type People interface {
	Disable()
}

type Client struct {
	ID     int
	Name   string
	Active bool
}

type Company struct {
	Name string
}

func (c *Company) Disable() {

}

func (c *Client) Disable() {

	c.Active = false
	fmt.Printf("O cliente %s foi desativado \n", c.Name)
}

func Disabling(people People) {
	people.Disable()
}

func main() {
	client := Client{
		ID:     1,
		Name:   "tau",
		Active: true,
	}
	fmt.Println(client)
	Disabling(&client)

	fmt.Println(client)

	company := Company{Name: "DC"}
	Disabling(&company)

	fmt.Println(company)

}
