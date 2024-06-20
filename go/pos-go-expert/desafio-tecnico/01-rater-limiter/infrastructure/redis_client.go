package infrastructure

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisRepositoryImpl struct {
	Conn *redis.Client
}

var ctx = context.Background()

func NewRedis() {
	//conn *gorm.DB) repository.NewsRepository
	//Conn *gorm.DB
	Conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//}

	//func AddMessage()
	//{
	err := Conn.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := Conn.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := Conn.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
