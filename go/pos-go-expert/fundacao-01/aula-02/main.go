package main

// declaracao e atribuicao de variavel

// variaveis globais
var a string = "Hello, World"
var (
	b bool
	c int
	d string
	e float64
)

var (
	b_atribuida bool    = true
	c_atribuida int     = 10
	d_atribuida string  = "Dev"
	e_atribuida float64 = 1.2
)

func main() {

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)

	// variavel local
	var a_local string
	print(a_local)

	println(b_atribuida)
	println(c_atribuida)
	println(d_atribuida)
	println(e_atribuida)

	// var nao_usada string // caso não seja utilizada,  nao sera possível executar o programa

	// declarando e setando valor na variavel
	// apenas na primeira iteracao, se eu tentar replicar essa rotina, ele vai informar erro
	var_setada := 10 // tipo inteiro

	println(var_setada)
}
