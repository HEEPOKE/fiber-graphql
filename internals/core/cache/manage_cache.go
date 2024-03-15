package cache

import (
	"context"
	"time"

	"github.com/HEEPOKE/fiber-graphql/pkg/database"
)

func SetKey(key string, value []byte) error {
	client := database.ConnectToRedis()
	expiration := 24 * time.Hour
	return client.Set(context.Background(), key, value, expiration).Err()
}

func GetKey(key string) (string, error) {
	client := database.ConnectToRedis()
	return client.Get(context.Background(), key).Result()
}
