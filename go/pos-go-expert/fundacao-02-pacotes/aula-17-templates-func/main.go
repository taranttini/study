package main

import (
	//esse html/template permite uma camada de seguranca para o modelo de paginas
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(text string) string {
	return strings.ToUpper(text)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"header.html",
			"template.html",
		}

		tmp := template.New("template.html")
		tmp.Funcs(template.FuncMap{"ToUpper": ToUpper})
		tmp = template.Must(tmp.ParseFiles(templates...))

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
