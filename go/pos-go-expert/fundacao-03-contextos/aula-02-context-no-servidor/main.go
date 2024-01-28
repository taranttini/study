package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		// no console print
		log.Println("Request processada com sucesso!")
		// print no navegador
		w.Write([]byte("Resquest processada com sucesso!!!"))

	case <-ctx.Done():
		// no console print
		log.Println("Request cancelada pelo cliente!")
	}

}
