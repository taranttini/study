package main

import "fmt"

// importando fmt e tipagem

type ID int

func main() {
	var dado ID = 10
	//println(dado)
	fmt.Printf("O tipo de \"dado\" é %T", dado)

	var dado_int int16
	fmt.Printf("\nO tipo de \"dado_int\" é %T", dado_int)
}
