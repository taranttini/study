package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmp := template.Must(
			template.New("template.html").ParseFiles("template.html"))

		err := tmp.Execute(w, Cursos{
			{"Go", 40},
			{"C#", 50},
			{"Java", 60},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)

}
