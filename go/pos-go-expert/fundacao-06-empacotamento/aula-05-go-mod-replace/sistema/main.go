package main

import "github.com/taranttini/study/go/pos-go-expert/fundacao-06-empacotamento/aula-05-go-mod-replace/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())

}
