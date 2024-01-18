package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	tmp := template.Must(
		template.New("template.txt").ParseFiles("template.txt"))

	err := tmp.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"C#", 50},
		{"Java", 60},
	})
	if err != nil {
		panic(err)
	}

}
