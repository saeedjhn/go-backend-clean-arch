package redisuser

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/cache/redis"
	"log"
)

type DB struct {
	conn redis.DB
}

func New(conn redis.DB) *DB {
	return &DB{conn: conn}
}

func (d *DB) Set() {
	log.Print("redis-user - IMPL ME")
}
