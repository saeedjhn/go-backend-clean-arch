package bootstrap

import (
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/cache/redis"
)

func NewRedisClient(config redis.Config) (redis.DB, error) {
	myDB := redis.New(config)

	if err := myDB.ConnectTo(); err != nil {
		return nil, err
	}

	return myDB, nil
}

func CloseRedisClient(redisClient redis.DB) error {
	if err := redisClient.Client().Close(); err != nil {
		return fmt.Errorf("don`t close redis client connection: %w", err)
	}

	return nil
}
