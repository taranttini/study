package math

// type Math struct -> acessivel ou exportado para uso externo
// type math struct -> nao acessivel, ou nao exportado para uso externo
type math struct {
	// A int -> acessível externamento
	// a int -> inacessível externamente
	a int
	b int
}

// func (m Math) Add() int -> acessível externalmente
// func (m Math) add() int -> inacessível externalmente
func (m math) Add() int {
	return m.a + m.b
}

func NewMath(a, b int) math {
	return math{a: a, b: b}
}
