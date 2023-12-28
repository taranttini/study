package main

import "fmt"

// percorrendo arrays

func main() {
	var meu_array [3]int
	meu_array[0] = 3
	meu_array[1] = 4
	meu_array[2] = 6

	fmt.Println(len(meu_array)) // tamanho do array

	fmt.Println(meu_array[len(meu_array)-1]) // valor do ultimo item no array

	for i, v := range meu_array {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}
