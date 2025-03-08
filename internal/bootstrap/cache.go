package bootstrap

import (
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/inmemory"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/redis"
)

func NewInMemory() *inmemory.DB {
	return inmemory.New()
}

func NewRedisClient(c redis.Config) (*redis.DB, error) {
	db := redis.New(c)

	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseRedisClient(db *redis.DB) error {
	if err := db.Client().Close(); err != nil {
		return fmt.Errorf("don`t close redis client connection: %w", err)
	}

	return nil
}
