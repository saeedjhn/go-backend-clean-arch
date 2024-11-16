package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	config Config
	db     *redis.Client
}

func New(config Config) *Redis {
	return &Redis{config: config}
}

func (r *Redis) ConnectTo() error {
	var err error

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.config.Host, r.config.Port),
		Password: r.config.Password,
		DB:       r.config.DB,
	})

	if err = rdb.Ping(context.Background()).Err(); err != nil {
		return err
	}
	r.db = rdb

	return nil
}

func (r *Redis) Client() *redis.Client {
	return r.db
}
