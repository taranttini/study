package main

import "fmt"

// generics

func sum_map(values map[string]int) int {
	var total int
	for _, value := range values {
		total += value
	}
	return total
}

func sum_map_generic[T int | float64](values map[string]T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

// criando uma interface para ser minha constraint generica
type number interface {
	int | float64
}

func sum_map_generic_constraint[T number](values map[string]T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

type my_number int

// criando uma interface para ser minha constraint generica
// como eu tenho o tipo ~int caso eu crie algum tipo especifico com int,
// ele acabara tambem aceitando, no caso do exemplo abaixo estarei usando o MyNumber
// como parametro ao chamar o metodo/funcao
type number_or_equivalent interface {
	~int | float64
}

func sum_map_generic_equivalent[T number_or_equivalent](values map[string]T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

func first_value_is_more_than_second_value[T number_or_equivalent](value1 T, value2 T) bool {
	fmt.Printf("%v is more than %v ? ", value1, value2)
	return value1 > value2
}

func main() {
	values := map[string]int{"a": 10, "b": 20, "c": 30}
	println(sum_map(values))

	values_float := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	fmt.Printf("sum_map_generic %v \n", sum_map_generic(values_float))

	fmt.Printf("sum_map_generic_constraint %v \n", sum_map_generic_constraint(values_float))

	values_myNumbers := map[string]my_number{"a": 1, "b": 2, "c": 3}

	fmt.Printf("sum_map_generic_equivalent %v \n", sum_map_generic_equivalent(values_myNumbers))

	println(first_value_is_more_than_second_value(3, 2.1))
}
