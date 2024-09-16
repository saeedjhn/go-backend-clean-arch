package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type DB interface {
	Client() *redis.Client
}

type RedisDB struct {
	config Config
	db     *redis.Client
	err    error
}

var _ DB = (*RedisDB)(nil)

func New(config Config) *RedisDB {
	return &RedisDB{config: config}
}

func (r *RedisDB) ConnectTo() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.config.Host, r.config.Port),
		Password: r.config.Password,
		DB:       r.config.DB,
	})

	if r.err = rdb.Ping(context.Background()).Err(); r.err != nil {
		return r.err
	}
	r.db = rdb

	return nil
}

func (r *RedisDB) Client() *redis.Client {
	return r.db
}
