package main

import (
	"fmt"
)

// closures

func main() {
	// funcao dentro de outra funcao, onde seu efeito e somente dentro desse contexto
	multipler_sum := func(numbers ...int) int {
		return sum(numbers...) * 2
	}(25, 25)

	func() {
		// funcacao anonima que executa alguma coisa antes
	}()

	fmt.Println(multipler_sum)

}

// funcao que recebe inumeros valores e soma todos eles
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
