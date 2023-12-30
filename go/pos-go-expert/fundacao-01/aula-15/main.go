package main

// ponteiros

func main() {

	// memoria -> endereco -> valor
	a := 10
	println(&a)

	var pointer *int = &a

	println(pointer)

	*pointer = 20

	println(a)

	b := &a

	println(b)

	println(*b) // dereference, aqui ele consegue exibir o valor de b

	*b = 30
	println(a)

}
