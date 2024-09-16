package redis

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	c := Config{
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

	redis := New(c)
	if err := redis.ConnectTo(); err != nil {
		log.Fatal(err)
	}

	log.Println(redis.Client().Conn())

	expiration := time.Duration(200 * float64(time.Second))
	log.Println(redis.Client().Set(context.Background(), "KEY", "VALUE", expiration))
}
