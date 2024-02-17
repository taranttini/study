package main

import (
	"github.com/google/uuid"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-06-empacotamento/aula-06-workspace/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
