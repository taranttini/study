package main

// ponteiros quando usar

// dessa forma a funcao recebe uma copia dos valores, e nao modifica os valores originais
func sum(value1, value2 int) int {
	value1 += 10

	return value1 + value2
}

// dessa forma estou passando o espaco de memoria das variais, e assim ela modificam o valor original
func sum_pointers(value1, value2 *int) int {
	*value1 += 10

	return *value1 + *value2
}

func main() {

	var_value1 := 10
	var_value2 := 10

	println(sum(var_value1, var_value2))
	println(var_value1)

	println("modificando o valor original")

	println(sum_pointers(&var_value1, &var_value2))
	println(var_value1)
}
