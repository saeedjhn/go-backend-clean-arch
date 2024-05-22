package bootstrap

import (
	postgresql2 "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/postgresql"
	"log"
)

func NewPostgresqlDB(config postgresql2.Config) postgresql2.DB {
	return postgresql2.New(config)
}

func ClosePostgresqlDB(postgresqlDB postgresql2.DB) {
	err := postgresqlDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close postgresql connection: %s", err.Error())
	}
}
