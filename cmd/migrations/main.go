package main

import (
	"go-backend-clean-arch/configs"
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/infrastructure/persistance/db/mysql/migratormysql"
	"go-backend-clean-arch/internal/infrastructure/persistance/db/pq/migratorpq"
)

func main() {
	app := bootstrap.App(configs.Development)

	// MysQL
	mysqlDir := "./internal/repository/migrations/mysqlmigration"
	migratorMysql := migratormysql.New(app.MysqlDB, mysqlDir)
	migratorMysql.Down()
	migratorMysql.Up()

	// Pq
	pqDir := "./internal/repository/migrations/pqmigration"
	migratorPq := migratorpq.New(app.PostgresDB, pqDir)
	migratorPq.Down()
	migratorPq.Up()

	// etc
}
