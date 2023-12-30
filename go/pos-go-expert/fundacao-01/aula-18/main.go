package main

import "fmt"

// interfaces vazias

func main() {
	var x interface{} = 10
	var y interface{} = "valor texto"

	showType(x)
	showType(y)
}

func showType(x interface{}) {
	fmt.Printf("tipo = %T, valor = %v \n", x, x)
}
