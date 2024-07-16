package usecase

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rate-limiter/configs"
	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rate-limiter/internal/entity"
)

type RateLimiterUseCase struct {
	RateLimiterRepository entity.RateLimiterRepositoryInterface
}

func NewRateLimiterUseCase(
	RateLimiterRepository entity.RateLimiterRepositoryInterface,
) *RateLimiterUseCase {
	return &RateLimiterUseCase{
		RateLimiterRepository: RateLimiterRepository,
	}
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

func (c *RateLimiterUseCase) _RequestIsInvalid(keyType string, keyValue string, qtyRequestAcceptable int, keepRequest int) bool {

	//config, _ := configs.LoadConfig(".")

	err := c.RateLimiterRepository.Insert(keyType, keyValue, keepRequest)
	if err != nil {
		panic(err)
	}
	qtyRequest, err := c.RateLimiterRepository.Count(keyType, keyValue)
	if err != nil {
		panic(err)
	}
	if qtyRequest > qtyRequestAcceptable {
		fmt.Printf("%s [ %s ] request limit over \n", keyType, keyValue)
		return true
	}
	fmt.Printf("%s [ %s ] has called [ %v ] time(s)\n", keyType, keyValue, qtyRequest)
	return false
}

func (c *RateLimiterUseCase) _RequestIsBlocked(keyType string, keyValue string) bool {
	qtyRequest, err := c.RateLimiterRepository.Count("BLOCK", fmt.Sprintf("%s-%s", keyType, keyValue))
	if err != nil {
		panic(err)
	}
	if qtyRequest > 0 {
		fmt.Printf("%s [ %s ] request limit over - BLOCK \n", "BLOCK", fmt.Sprintf("%s-%s", keyType, keyValue))
		return true
	}

	return false
}

func (c *RateLimiterUseCase) _BlockItem(keyType string, keyValue string, blockTime int) {
	c.RateLimiterRepository.Insert("BLOCK", fmt.Sprintf("%s-%s", keyType, keyValue), blockTime)
}

func (c *RateLimiterUseCase) RequestLimitEnd(w http.ResponseWriter, r *http.Request) bool {

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

	if c._RequestIsBlocked(keyType, keyValue) {
		//ResearchLimitHasEnd(w)
		return true
	}

	if c._RequestIsInvalid(keyType, keyValue, requestQty, config.KEEP_REQUEST_PER_X_SECONDS) {
		c._BlockItem(keyType, keyValue, blockedTime)
		//c.ResearchLimitHasEnd(w)
		return true
	}

	return false
}
