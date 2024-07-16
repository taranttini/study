package database

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"

	"github.com/redis/go-redis/v9"
)

// Test Redis is down.
func Test_RedisWith2Messages(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := NewRateLimiterRepository(fakeRedisClient)

	fakeDb.Insert("IP", "1.0.0.1", 10)
	fakeDb.Insert("IP", "1.0.0.1", 10)
	count, err := fakeDb.Count("IP", "1.0.0.1")

	assert.NoError(t, err)
	assert.Equal(t, count, 2)
	//assert.Nil(t, err)
}

func Test_RedisWith0Messages(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := NewRateLimiterRepository(fakeRedisClient)

	fakeDb.Insert("IP", "1.0.0.1", 10)
	fakeDb.Insert("IP", "1.0.0.1", 10)

	// avancando no tempo pra limpar a fila
	mr.FastForward(11 * time.Second)

	count, err := fakeDb.Count("IP", "1.0.0.1")

	assert.NoError(t, err)
	assert.Equal(t, count, 0)
	//assert.Nil(t, err)
}
