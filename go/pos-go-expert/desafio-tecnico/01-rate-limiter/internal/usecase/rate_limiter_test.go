package usecase

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/01-rate-limiter/infra/database"
)

func Test_RequestAreValid(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := database.NewRateLimiterRepository(fakeRedisClient)
	rateLimiterUseCase := NewRateLimiterUseCase(fakeDb)

	requestQtyLimit := 10
	requestExpires := 10 //seconds
	expireTime := 10     //seconds

	fakeDb.Insert("IP", "1.0.0.1", requestExpires)
	fakeDb.Insert("IP", "1.0.0.1", requestExpires)

	fakeDb.Insert("IP", "1.0.0.2", requestExpires)
	fakeDb.Insert("IP", "1.0.0.2", requestExpires)

	result := rateLimiterUseCase._RequestIsInvalid("IP", "1.0.0.1", requestQtyLimit, expireTime)
	assert.False(t, result)

	result = rateLimiterUseCase._RequestIsInvalid("IP", "1.0.0.2", requestQtyLimit, expireTime)

	assert.False(t, result)
}

func Test_LimitAndBlocks(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := database.NewRateLimiterRepository(fakeRedisClient)
	rateLimiterUseCase := NewRateLimiterUseCase(fakeDb)

	requestQtyLimit := 3
	expireTime := 10
	blockTime := 20

	// validar limite sem nenhum item bloqueado ou inserido
	result := rateLimiterUseCase._CheckLimitEnd("IP", "1.0.0.1", requestQtyLimit, expireTime, blockTime)
	assert.False(t, result)

	result = rateLimiterUseCase._CheckLimitEnd("IP", "1.0.0.1", requestQtyLimit, expireTime, blockTime)
	assert.False(t, result)

	qtd, _ := rateLimiterUseCase.RateLimiterRepository.Count("IP", "1.0.0.1")
	assert.Equal(t, qtd, 2)

	result = rateLimiterUseCase._CheckLimitEnd("IP", "1.0.0.1", requestQtyLimit, expireTime, blockTime)
	assert.False(t, result)

	qtd, _ = rateLimiterUseCase.RateLimiterRepository.Count("IP", "1.0.0.1")
	assert.Equal(t, qtd, 3)

	result = rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")
	assert.False(t, result)

	qtd, _ = rateLimiterUseCase.RateLimiterRepository.Count("IP", "1.0.0.1")
	assert.Equal(t, qtd, 3)

	result = rateLimiterUseCase._CheckLimitEnd("IP", "1.0.0.1", requestQtyLimit, expireTime, blockTime)
	assert.True(t, result)

	qtd, _ = rateLimiterUseCase.RateLimiterRepository.Count("IP", "1.0.0.1")
	assert.Equal(t, qtd, 4)

	result = rateLimiterUseCase._CheckLimitEnd("IP", "1.0.0.1", requestQtyLimit, expireTime, blockTime)
	assert.True(t, result)

	qtd, _ = rateLimiterUseCase.RateLimiterRepository.Count("IP", "1.0.0.1")
	assert.Equal(t, qtd, 4)

	result = rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")
	assert.True(t, result)

	// avancando no tempo 10s pra limpar a fila
	mr.FastForward(10 * time.Second)

	qtd, _ = rateLimiterUseCase.RateLimiterRepository.Count("IP", "1.0.0.1")
	assert.Equal(t, qtd, 0)

	result = rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")
	assert.True(t, result)

	result = rateLimiterUseCase._CheckLimitEnd("IP", "1.0.0.1", requestQtyLimit, expireTime, blockTime)
	assert.True(t, result)

	qtd, _ = rateLimiterUseCase.RateLimiterRepository.Count("IP", "1.0.0.1")
	assert.Equal(t, qtd, 0)

	// avancando no tempo 10s pra limpar a fila
	mr.FastForward(10 * time.Second)

	result = rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")
	assert.False(t, result)

	result = rateLimiterUseCase._CheckLimitEnd("IP", "1.0.0.1", requestQtyLimit, expireTime, blockTime)
	assert.False(t, result)

	qtd, _ = rateLimiterUseCase.RateLimiterRepository.Count("IP", "1.0.0.1")
	assert.Equal(t, qtd, 1)

}

func Test_RequestAreInvalidAndOtherNot(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := database.NewRateLimiterRepository(fakeRedisClient)
	rateLimiterUseCase := NewRateLimiterUseCase(fakeDb)

	requestQtyLimit := 4
	expireTime := 10

	fakeDb.Insert("IP", "1.0.0.1", expireTime)
	fakeDb.Insert("IP", "1.0.0.1", expireTime)
	fakeDb.Insert("IP", "1.0.0.1", expireTime)
	fakeDb.Insert("IP", "1.0.0.1", expireTime)
	fakeDb.Insert("IP", "1.0.0.1", expireTime)

	fakeDb.Insert("IP", "1.0.0.2", expireTime)
	fakeDb.Insert("IP", "1.0.0.2", expireTime)

	result := rateLimiterUseCase._RequestIsInvalid("IP", "1.0.0.1", requestQtyLimit, expireTime)
	assert.True(t, result)

	result = rateLimiterUseCase._RequestIsInvalid("IP", "1.0.0.2", requestQtyLimit, expireTime)
	assert.False(t, result)
}

func Test_IpAreBlockedAndCheckValid(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := database.NewRateLimiterRepository(fakeRedisClient)
	rateLimiterUseCase := NewRateLimiterUseCase(fakeDb)

	expireTime := 20

	rateLimiterUseCase._BlockItem("IP", "1.0.0.1", expireTime)

	result := rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")
	assert.True(t, result)

	// avancando no tempo 10s pra limpar a fila
	mr.FastForward(10 * time.Second)

	result = rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")
	assert.True(t, result)

	// avancando no tempo 10s pra limpar a fila
	mr.FastForward(10 * time.Second)

	result = rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")
	assert.False(t, result)
}

func Test_IpAreUnblocked(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := database.NewRateLimiterRepository(fakeRedisClient)
	rateLimiterUseCase := NewRateLimiterUseCase(fakeDb)

	expireTime := 10

	fakeDb.Insert("IP", "1.0.0.1", expireTime)

	result := rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")

	assert.False(t, result)
}
