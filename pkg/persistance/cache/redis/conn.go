package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	ctx    context.Context
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

	if err = rdb.Ping(r.checkCtx()).Err(); err != nil {
		return err
	}
	r.db = rdb

	return nil
}

func (r *DB) SetCtx(ctx context.Context) *DB {
	r.ctx = ctx

	return r
}

func (r *DB) Ping() error {
	return r.db.Ping(r.checkCtx()).Err()
}

func (r *DB) Client() *redis.Client {
	return r.db
}

func (r *DB) checkCtx() context.Context {
	if r.ctx == nil {
		r.ctx = context.Background()
	}

	return r.ctx
}
