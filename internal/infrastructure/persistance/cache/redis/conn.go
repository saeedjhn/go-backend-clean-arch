package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type DB interface {
	Client() *redis.Client
}

type Redis struct {
	config Config
	db     *redis.Client
	err    error
}

var _ DB = (*Redis)(nil)

func New(config Config) *Redis {
	return &Redis{config: config}
}

func (r *Redis) ConnectTo() error {
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

func (r *Redis) Client() *redis.Client {
	return r.db
}
