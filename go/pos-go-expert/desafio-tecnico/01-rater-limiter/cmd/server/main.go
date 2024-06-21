package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rater-limiter/configs"

	redisClient "github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rater-limiter/infrastructure"
)

// var sl []string
var rd = redisClient.NewRedis()

func main() {
	_, err := configs.LoadConfig("../../")
	if err != nil {
		panic(err)
	}

	redisClient.Insert(rd, "chave1", "1", 15)
	redisClient.Insert(rd, "chave1", "1", 20)
	redisClient.Insert(rd, "chave1", "1", 25)
	redisClient.Insert(rd, "chave2", "2", 10)
	redisClient.Insert(rd, "chave2", "3", 30)
	redisClient.Count(rd, "*", "*")
	//redisClient.GetMessage(rd, "chave2")

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
		//log.Println("_METHOD_", r.Method, "_PATH_", r.URL.Path, r.Host, "_IP_", GetUserIp(r))

		if RequestLimitEnd(w, r) {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RequestLimitEnd(w http.ResponseWriter, r *http.Request) bool {

	var keyType string
	var keyValue string
	var requestQty int
	var blockedTime int

	tk := GetToken(r)
	config, _ := configs.LoadConfig(".")

	if tk != "" {
		keyType = "TOKEN"
		keyValue = tk
		requestQty = config.QTY_REQUEST_TOKEN
		blockedTime = config.BLOCKED_TOKEN_PER_X_SECONDS
	} else {
		keyType = "IP"
		keyValue = GetUserIp(r)
		requestQty = config.QTY_REQUEST_IP
		blockedTime = config.BLOCKED_IP_PER_X_SECONDS
	}

	if RequestIsBlocked(keyType, keyValue) {
		ResearchLimitHasEnd(w)
		return true
	}

	if RequestIsValid(keyType, keyValue, requestQty) {
		BlockItem(keyType, keyValue, blockedTime)
		ResearchLimitHasEnd(w)
		return true
	}

	return false
}

func ResearchLimitHasEnd(w http.ResponseWriter) {
	w.WriteHeader(http.StatusTooManyRequests)
	json.NewEncoder(w).Encode("you have reached the maximum number of requests or actions allowed within a certain time frame")
}

func BlockItem(keyType string, keyValue string, blockTime int) {
	redisClient.Insert(rd, "BLOCK", fmt.Sprintf("%s-%s", keyType, keyValue), blockTime)
}

func RequestIsBlocked(keyType string, keyValue string) bool {
	qtyRequest := redisClient.Count(rd, "BLOCK", fmt.Sprintf("%s-%s", keyType, keyValue))

	if qtyRequest > 0 {
		fmt.Printf("%s [ %s ] request limit over - BLOCK \n", "BLOCK", fmt.Sprintf("%s-%s", keyType, keyValue))
		return true
	}
	return false
}

func RequestIsValid(keyType string, keyValue string, qtyRequestAcceptable int) bool {

	config, _ := configs.LoadConfig(".")

	redisClient.Insert(rd, keyType, keyValue, config.KEEP_REQUEST_PER_X_SECONDS)
	qtyRequest := redisClient.Count(rd, keyType, keyValue)

	if qtyRequest > qtyRequestAcceptable {
		fmt.Printf("%s [ %s ] request limit over \n", keyType, keyValue)
		return true
	}
	fmt.Printf("%s [ %s ] has called [ %v ] time(s)\n", keyType, keyValue, qtyRequest)
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
