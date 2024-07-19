package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rate-limiter/configs"
	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rate-limiter/infra/database"
	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rate-limiter/internal/usecase"

	"github.com/redis/go-redis/v9"
)

var rateLimiterUseCase *usecase.RateLimiterUseCase

func main() {

	fmt.Print("##### DOCKER step 1 \n")

	_, err := configs.LoadConfig("../../")
	if err != nil {
		panic(err)
	}

	fmt.Print("##### DOCKER step 2 \n")

	db := *redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		//Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Print("##### DOCKER step 3 \n")

	rateLimiterUseCase = usecase.NewRateLimiterUseCase(
		database.NewRateLimiterRepository(&db),
	)

	// database.NewRateLimiterRepository(&db).Insert("chave1", "1", 15)
	// database.NewRateLimiterRepository(&db).Insert("chave1", "1", 20)
	// database.NewRateLimiterRepository(&db).Insert("chave1", "1", 25)
	// database.NewRateLimiterRepository(&db).Insert("chave2", "2", 10)
	// database.NewRateLimiterRepository(&db).Insert("chave2", "3", 30)
	// database.NewRateLimiterRepository(&db).Count("*", "*")

	fmt.Print("runing api go Rate Limiter \n")

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
		//log.Println("_METHOD_", r.Method, "_PATH_", r.URL.Path, r.Host, "_IP_", GetUserIp(r))

		if rateLimiterUseCase.RequestLimitEnd(w, r) {
			ResearchLimitHasEnd(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ResearchLimitHasEnd(w http.ResponseWriter) {
	w.WriteHeader(http.StatusTooManyRequests)
	json.NewEncoder(w).Encode("you have reached the maximum number of requests or actions allowed within a certain time frame")
}
