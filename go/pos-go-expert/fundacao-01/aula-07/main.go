package main

import "fmt"

// maps
// funciona como chave=valor, chave=valor
// e ordena as chave
func main() {
	salary := map[string]int{"tau": 100, "benjamin": 500, "nicole": 500}
	fmt.Println(salary)        // listar todos
	fmt.Println(salary["tau"]) // exibir especifico
	delete(salary, "tau")      // remover especifico
	fmt.Println(salary)
	salary["novo"] = 150 // incluir novo item
	fmt.Println(salary)

	salary2 := make(map[string]int) // criando lista vazia chave string, valor int
	salary3 := map[int]string{}     // criando lista vazia cahve int, valor string

	fmt.Println(salary2)
	fmt.Println(salary3)

	// percorrendo todos os itens da lista e usando a chave e valor
	for key, value := range salary {
		fmt.Printf("%s = %d \n", key, value)
	}

	// percorrendo todos os itens da lista e usando somente o valor
	// o uso do _ significa blank identifier
	for _, value := range salary {
		fmt.Println(value)
	}
}
