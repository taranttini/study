package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RateLimiterRepository struct {
	DB *redis.Client
}

var ctx = context.Background()

func NewRateLimiterRepository(db *redis.Client) *RateLimiterRepository {
	return &RateLimiterRepository{
		DB: db,
	}
	/*redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	*/
}

func (r *RateLimiterRepository) Insert(rateType string, value string, expiresTimeInSecond int) error {
	key := fmt.Sprintf("%s:%s-%s", rateType, value, uuid.New())
	fmt.Print("ITEM ", key, " DURAR ", expiresTimeInSecond, "\n")
	err := r.DB.Set(ctx, key, "1", time.Duration(expiresTimeInSecond)*time.Second).Err()
	if err != nil {
		return err //panic(err)
	}
	return nil
}

func (r *RateLimiterRepository) Count(rateType string, value string) (int, error) {
	key := fmt.Sprintf("%s:%s-*", rateType, value)
	val, err := r.DB.Keys(ctx, key).Result()
	if err != nil {
		return 0, err //panic(err)
	}

	return len(val), nil
}
