package main

// loop

func main() {
	// no go existe apenas o for

	for i := 0; i < 10; i++ {
		println(i)
	}

	// funciona como se fosse um foreach
	numeros := []string{"um", "dois", "tres"}
	for indice, _ := range numeros {
		println(indice)
	}

	// funciona como se fosse um foreach
	for _, valor := range numeros {
		println(valor)
	}

	// funciona como se fosse um while
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// exemplo de loop infinito
	// for { println("looping infinito") }

}
