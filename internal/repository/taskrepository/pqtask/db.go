package pqtask

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/pq"
	"log"
)

type DB struct {
	conn pq.DB
}

func New(conn pq.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) List() {
	log.Print("postgres-taskgateway -> Find - IMPL ME")
}
