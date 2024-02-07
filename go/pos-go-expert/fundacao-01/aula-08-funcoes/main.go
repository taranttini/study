package main

import (
	"errors"
	"fmt"
)

// funcoes

func main() {
	fmt.Println(sum(10, 5))
	fmt.Println(sum2(10, 5))
	fmt.Println(sum3(10, 15))
	fmt.Println(sum_check_error(1, 1))
	fmt.Println(sum_check_error(10, 1))
}

// funcao recebe o nome da varivel e o tipo, alem do tipo de saida
func sum(value1 int, value2 int) int {
	return value1 + value2
}

// quando o tipo de entrar for igual para os parametros, pode passar 1 so vez
func sum2(value1, value2 int) int {
	return value1 + value2
}

// funcao com 2 saidas, uma com valor, e outra validando verdadeiro ou falso
func sum3(value1, value2 int) (int, bool) {
	calc := value1 + value2
	return calc, calc > 20
}

// funcao com saida de erro ou sem erro
func sum_check_error(value1, value2 int) (int, error) {
	calc := value1 + value2
	if calc > 10 {
		return calc, errors.New("soma superior a 10")
	}
	return calc, nil
}
