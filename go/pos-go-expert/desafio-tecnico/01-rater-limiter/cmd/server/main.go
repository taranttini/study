package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rater-limiter/configs"
)

var sl []string

func main() {
	_, err := configs.LoadConfig("../../")
	if err != nil {
		panic(err)
	}

	fmt.Print("api go\n")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(LogRequest)

	r.Route("/", func(r chi.Router) {
		r.Get("/", Rota)
		r.Get("/page-1", Rota)
		r.Get("/page-2", Rota)
	})

	http.ListenAndServe(":8080", r)

}

func Rota(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Request to %s", r.URL.Path))

}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("_METHOD_", r.Method, "_PATH_", r.URL.Path, r.Host, "_IP_", GetUserIp(r))

		if RequestLimitEnd(w, r) {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RequestLimitEnd(w http.ResponseWriter, r *http.Request) bool {
	tk := GetToken(r)

	config, _ := configs.LoadConfig(".")

	if tk != "" {

		existsToken := RequestValid("TOKEN", tk, config.QTY_REQUEST_TOKEN)
		if existsToken {
			ResearchLimitHasEnd(w)
			return true
		}
	} else {
		ip := GetUserIp(r)
		exists := RequestValid("IP", ip, config.QTY_REQUEST_IP)
		if exists {
			ResearchLimitHasEnd(w)
			return true
		}
	}
	return false
}

func ResearchLimitHasEnd(w http.ResponseWriter) {
	w.WriteHeader(http.StatusTooManyRequests)
	json.NewEncoder(w).Encode("you have reached the maximum number of requests or actions allowed within a certain time frame")
}

func RequestValid(_type string, value string, limit int) bool {
	sl = append(sl, value)

	counter := 0
	for _, num := range sl {
		if num == value {
			counter++
		}
	}
	if counter > limit {
		fmt.Printf("%s [ %s ] request limit over \n", _type, value)
		return true
	}
	fmt.Printf("%s [ %s ] has called [ %v ] time(s)\n", _type, value, counter)
	return false
}

func GetToken(r *http.Request) string {
	API_TOKEN := r.Header.Get("API_TOKEN")

	if API_TOKEN == "" {
		return ""
	}
	return API_TOKEN
}

func GetUserIp(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-IP")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	if strings.Contains(IPAddress, ":") {
		IPAddress = strings.Split(IPAddress, ":")[0]
	}
	return IPAddress
}
