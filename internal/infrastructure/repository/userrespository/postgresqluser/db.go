package postgresqluser

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/postgresql"
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

func (r *DB) Create() {
	log.Print("mysql-user -> Create - IMPL ME")
}
