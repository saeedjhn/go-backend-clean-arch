package bootstrap

import (
	mysql2 "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/mysql"
	"log"
)

func NewMysqlDB(config mysql2.Config) mysql2.DB {
	return mysql2.New(config)
}

func CloseMysqlDB(mysqlDB mysql2.DB) {
	err := mysqlDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close mysql connection: %s", err.Error())
	}
}
