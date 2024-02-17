package main

import (
	"fmt"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-06-empacotamento/aula-03-exportar-objetos/math"
)

func main() {

	//erro nao acessível
	// m := math.math{a: 1, B: 2}
	// fmt.Println(m.Add(1, 2))

	// acessivel
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
	//fmt.Println(math.X)
}
