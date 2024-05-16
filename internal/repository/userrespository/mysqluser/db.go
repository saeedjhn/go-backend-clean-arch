package mysqluser

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/db/mysql"
	"log"
)

type DB struct {
	conn mysql.DB
}

func New(conn mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (r *DB) Create() {
	log.Print("mysql-user -> Create - IMPL ME")
}
