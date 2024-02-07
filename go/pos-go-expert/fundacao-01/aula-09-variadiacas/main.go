package main

import (
	"fmt"
)

// funcoes variadicas

func main() {
	fmt.Println(sum(10, 5, 5, 10, 20))

}

// funcao que recebe inumeros valores e soma todos eles
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
