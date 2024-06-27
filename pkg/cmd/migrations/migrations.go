package migrations

import (
	"flag"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql/migratormysql"
	"log"
)

const (
	mysqlDIR = "./internal/repository/migrations/mysqlmigration"
	pqDir    = "./internal/repository/migrations/pqmigration"
)

const (
	ArgUP       = "up"
	ArgDown     = "down"
	ArgRollback = "rollback"
)

func Start(app *bootstrap.Application) {
	// Define flags
	up := flag.Bool(ArgUP, false, "Example flag to demonstrate passing --up")
	down := flag.Bool(ArgDown, false, "Example flag to demonstrate passing --down")
	rollback := flag.Bool(ArgRollback, true, "Example flag to demonstrate passing --down")

	// Parse flags
	flag.Parse()

	log.Println("Startup migrations...")

	HandleFlag(app, *up, *down, *rollback)

	log.Println("finished migrations.")
}

func HandleFlag(app *bootstrap.Application, up, down, rollback bool) {
	switch {
	case up:
		Up(app)
	case down:
		Down(app)
	case rollback:
		Rollback(app)
	default:
		Rollback(app)
	}
}

func Up(app *bootstrap.Application) {
	// Mysql
	migratorMysql := migratormysql.New(app.MysqlDB, mysqlDIR)
	migratorMysql.Up()

	// Pq
	//migratorPq := migratorpq.New(app.PostgresDB, pqDIR)
	//migratorPq.Up()

	// Etc
}

func Down(app *bootstrap.Application) {
	// Mysql
	migratorMysql := migratormysql.New(app.MysqlDB, mysqlDIR)
	migratorMysql.Down()

	// Pq
	//migratorPq := migratorpq.New(app.PostgresDB, pqDIR)
	//migratorPq.Down()

	// Etc
}

func Rollback(app *bootstrap.Application) {

	// Mysql
	migratorMysql := migratormysql.New(app.MysqlDB, mysqlDIR)
	migratorMysql.Down()
	migratorMysql.Up()

	// Pq
	//migratorPq := migratorpq.New(app.PostgresDB, pqDIR)
	//migratorPq.Down()
	//migratorPq.Up()

	// etc
}
