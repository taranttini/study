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

func Test_RedisLimitMessagesCount(t *testing.T) {

	mr, _ := miniredis.Run()
	defer mr.Close()
	fakeRedisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	fakeDb := NewRateLimiterRepository(fakeRedisClient)

	fakeDb.Insert("IP", "1.0.0.1", 60)
	fakeDb.Insert("IP", "1.0.0.1", 60)
	fakeDb.Insert("IP", "1.0.0.1", 60)
	fakeDb.Insert("IP", "1.0.0.1", 30)
	fakeDb.Insert("IP", "1.0.0.1", 30)

	fakeDb.Insert("IP", "1.0.0.2", 60)
	fakeDb.Insert("IP", "1.0.0.2", 60)
	fakeDb.Insert("IP", "1.0.0.2", 30)
	fakeDb.Insert("IP", "1.0.0.2", 30)

	fakeDb.Insert("TOKEN", "XXX1", 60)
	fakeDb.Insert("TOKEN", "XXX1", 20)
	fakeDb.Insert("TOKEN", "XXX1", 20)

	fakeDb.Insert("TOKEN", "XXX2", 60)
	fakeDb.Insert("TOKEN", "XXX2", 60)
	fakeDb.Insert("TOKEN", "XXX2", 60)
	fakeDb.Insert("TOKEN", "XXX2", 60)
	fakeDb.Insert("TOKEN", "XXX2", 30)
	fakeDb.Insert("TOKEN", "XXX2", 30)

	count_ip1, err1 := fakeDb.Count("IP", "1.0.0.1")
	count_ip2, err2 := fakeDb.Count("IP", "1.0.0.2")
	count_tk1, err3 := fakeDb.Count("TOKEN", "XXX1")
	count_tk2, err4 := fakeDb.Count("TOKEN", "XXX2")

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NoError(t, err3)
	assert.NoError(t, err4)

	assert.Equal(t, count_ip1, 5)
	assert.Equal(t, count_ip2, 4)
	assert.Equal(t, count_tk1, 3)
	assert.Equal(t, count_tk2, 6)

	// avancando no tempo pra limpar a fila
	mr.FastForward(31 * time.Second)

	count_ip1, _ = fakeDb.Count("IP", "1.0.0.1")
	count_ip2, _ = fakeDb.Count("IP", "1.0.0.2")
	count_tk1, _ = fakeDb.Count("TOKEN", "XXX1")
	count_tk2, _ = fakeDb.Count("TOKEN", "XXX2")

	assert.Equal(t, count_ip1, 3)
	assert.Equal(t, count_ip2, 2)
	assert.Equal(t, count_tk1, 1)
	assert.Equal(t, count_tk2, 4)

	// avancando no tempo pra limpar a fila
	mr.FastForward(30 * time.Second)

	count_ip1, _ = fakeDb.Count("IP", "1.0.0.1")
	count_ip2, _ = fakeDb.Count("IP", "1.0.0.1")
	count_tk1, _ = fakeDb.Count("TOKEN", "XXX1")
	count_tk2, _ = fakeDb.Count("TOKEN", "XXX2")

	assert.Equal(t, count_ip1, 0)
	assert.Equal(t, count_ip2, 0)
	assert.Equal(t, count_tk1, 0)
	assert.Equal(t, count_tk2, 0)
}
