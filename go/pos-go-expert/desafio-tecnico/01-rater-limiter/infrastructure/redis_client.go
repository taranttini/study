package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisRepositoryImpl struct {
	Conn *redis.Client
}

var ctx = context.Background()

func NewRedis() *redis.Client {
	//conn *gorm.DB) repository.NewsRepository
	//Conn *gorm.DB
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Insert(cli *redis.Client, _type string, value string, expiresTimeInSecond int) {
	key := fmt.Sprintf("%s:%s-%s", _type, value, uuid.New())
	fmt.Print("ITEM ", key, " DURAR ", expiresTimeInSecond, "\n")
	err := cli.Set(ctx, key, "1", time.Duration(expiresTimeInSecond)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func Count(cli *redis.Client, _type string, value string) int {
	key := fmt.Sprintf("%s:%s-*", _type, value)
	val, err := cli.Keys(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return len(val)
}
