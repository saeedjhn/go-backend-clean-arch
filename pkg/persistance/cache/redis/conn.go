package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	config Config
	db     *redis.Client
}

func New(config Config) *DB {
	return &DB{config: config}
}

func (r *DB) ConnectTo() error {
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

func (r *DB) Client() *redis.Client {
	return r.db
}
