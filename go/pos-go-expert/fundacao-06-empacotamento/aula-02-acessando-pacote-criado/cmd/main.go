package main

import (
	"fmt"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-06-empacotamento/aula-02-acessando-pacote-criado/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
