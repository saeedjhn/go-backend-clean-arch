package redis_test

import (
	"context"
	redis2 "github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/redis"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	c := redis2.Config{
		Host:               "127.0.0.1",
		Port:               "6379",
		Password:           "123456",
		DB:                 0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		PoolTimeout:        0,
		IdleCheckFrequency: 0,
	}

	redis := redis2.New(c)
	if err := redis.ConnectTo(); err != nil {
		t.Fatal(err)
	}

	t.Log(redis.Client().Conn())

	expiration := time.Duration(200 * float64(time.Second))
	t.Log(redis.Client().Set(context.Background(), "KEY", "VALUE", expiration))
}