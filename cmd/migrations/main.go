package main

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql/migratormysql"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/pq/migratorpq"
	"log"
)

func main() {
	app := bootstrap.App(configs.Development)
	log.Printf("%#v", app)

	log.Println("Startup migrations...")

	// Mysql
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

	log.Println("finished migrations.")
}
