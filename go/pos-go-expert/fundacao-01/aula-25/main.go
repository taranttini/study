package main

// compilando

func main() {
	// build nativo do seu ambiente
	println("go build main.go")

	//
	// build nativo para windows
	println("GOOS=windows go build main.go")

	//
	// build nativo para mac
	println("GOOS=darwin go build main.go")

	//
	// build nativo para linux
	println("GOOS=linux go build main.go")
}
