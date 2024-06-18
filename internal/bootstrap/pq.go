package bootstrap

import (
	"go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
	"log"
)

func NewPostgresConnection(config pq.Config) pq.DB {
	return pq.New(config)
}

func ClosePostgresConnection(postgresDB pq.DB) {
	err := postgresDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close pq connection: %s", err.Error())
	}
}
