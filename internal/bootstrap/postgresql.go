package bootstrap

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/postgresql"
	"log"
)

func NewPostgresqlDB(config postgresql.Config) postgresql.DB {
	return postgresql.New(config)
}

func ClosePostgresqlDB(postgresqlDB postgresql.DB) {
	err := postgresqlDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close postgresql connection: %s", err.Error())
	}
}
