package postgresqltask

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/db/postgresql"
	"log"
)

type DB struct {
	conn postgresql.DB
}

func New(conn postgresql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Find() {
	log.Print("mysql-user -> Find - IMPL ME")
}
