package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type DB interface {
	Client() *redis.Client
}

type RedisDB struct {
	config Config
	db     *redis.Client
}

func New(config Config) *RedisDB {
	return (&RedisDB{config: config}).conn()
}

func (r *RedisDB) conn() *RedisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.config.Host, r.config.Port),
		Password: r.config.Password,
		DB:       r.config.DB,
	})
	r.db = rdb

	return r
}

func (r *RedisDB) Client() *redis.Client {
	return r.db
}
