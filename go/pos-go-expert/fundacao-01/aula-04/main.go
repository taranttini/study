package main

import "fmt"

// importando fmt e tipagem

type ID int

func main() {
	var data ID = 10
	//println(data)
	fmt.Printf("O tipo de \"data\" é %T", data)

	var data_int int16
	fmt.Printf("\nO tipo de \"data_int\" é %T", data_int)
}
