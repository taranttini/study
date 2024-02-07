package main

import "fmt"

// percorrendo arrays

func main() {
	var my_array [3]int
	my_array[0] = 3
	my_array[1] = 4
	my_array[2] = 6

	fmt.Println(len(my_array)) // tamanho do array

	fmt.Println(my_array[len(my_array)-1]) // valor do ultimo item no array

	for i, v := range my_array {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}
