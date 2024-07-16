package usecase

import (
	"testing"

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

	fakeDb.Insert("IP", "1.0.0.1", 10)
	fakeDb.Insert("IP", "1.0.0.1", 10)

	result := rateLimiterUseCase._RequestIsInvalid("IP", "1.0.0.1", 10, 10)

	assert.False(t, result)
}

func Test_RequestAreInvalid(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := database.NewRateLimiterRepository(fakeRedisClient)
	rateLimiterUseCase := NewRateLimiterUseCase(fakeDb)

	fakeDb.Insert("IP", "1.0.0.1", 10)
	fakeDb.Insert("IP", "1.0.0.1", 10)
	fakeDb.Insert("IP", "1.0.0.1", 10)
	fakeDb.Insert("IP", "1.0.0.1", 10)
	fakeDb.Insert("IP", "1.0.0.1", 10)

	limitOfRequest := 4
	result := rateLimiterUseCase._RequestIsInvalid("IP", "1.0.0.1", limitOfRequest, 10)

	assert.True(t, result)
}

func Test_IpAreBlocked(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := database.NewRateLimiterRepository(fakeRedisClient)
	rateLimiterUseCase := NewRateLimiterUseCase(fakeDb)

	//fakeDb.Insert("BLOCK", "IP-1.0.0.1", 10)
	rateLimiterUseCase._BlockItem("IP", "1.0.0.1", 10)

	result := rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")

	assert.True(t, result)
}

func Test_IpAreUnblocked(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := database.NewRateLimiterRepository(fakeRedisClient)
	rateLimiterUseCase := NewRateLimiterUseCase(fakeDb)

	fakeDb.Insert("IP", "1.0.0.1", 10)

	result := rateLimiterUseCase._RequestIsBlocked("IP", "1.0.0.1")

	assert.False(t, result)
}
