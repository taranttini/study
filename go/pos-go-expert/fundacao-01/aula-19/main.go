package main

import "fmt"

// type assertion

func main() {

	var value interface{} = "valor texto"

	println(value)
	println(value.(string))

	res, ok := value.(int) // funciona como validador do tipo

	// res seria o resultado da saida, e o ok dizendo se foi ou nao possivel converter o valor
	fmt.Printf("valor %v, foi possivel converter? %v\n", res, ok)

	// nova tentativa
	var value2 interface{} = 10
	res, ok = value2.(int)
	fmt.Printf("valor %v, foi possivel converter? %v\n", res, ok)
}
