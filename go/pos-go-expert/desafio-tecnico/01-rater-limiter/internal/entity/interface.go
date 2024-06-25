package entity

type RateLimiterRepositoryInterface interface {
	Insert(keyType string, keyValue string, expiresTimeInSecond int) error
	Count(keyType string, keyValue string) (int, error)
}
