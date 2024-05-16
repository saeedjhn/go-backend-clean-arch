package mysqltask

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

func (r *DB) Find() {
	log.Print("mysql-user -> Find - IMPL ME")
}