package bootstrap

import (
	"fmt"
	redis2 "github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/redis"
)

func NewRedisClient(c redis2.Config) (redis2.DB, error) {
	db := redis2.New(c)

	if err := db.ConnectTo(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseRedisClient(db redis2.DB) error {
	if err := db.Client().Close(); err != nil {
		return fmt.Errorf("don`t close redis client connection: %w", err)
	}

	return nil
}
