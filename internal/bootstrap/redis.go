package bootstrap

import (
	redis2 "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/cache/redis"
	"log"
)

func NewRedisClient(config redis2.Config) redis2.DB {
	return redis2.New(config)
}

func CloseRedisClient(redisClient redis2.DB) {
	err := redisClient.Client().Close()

	if err != nil {
		log.Fatalf("don`t close redis client connection: %s", err.Error())
	}
}
