package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/cache/redis"
	"log"
)

func NewRedisClient(config redis.Config) redis.DB {
	return redis.New(config)
}

func CloseRedisClient(redisClient redis.DB) {
	err := redisClient.Client().Close()

	if err != nil {
		log.Fatalf("don`t close redis client connection: %s", err.Error())
	}
}
