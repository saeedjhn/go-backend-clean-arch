package bootstrap

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/configs"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/cache/redis"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/mysql"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/postgresql"
)

type Application struct {
	Config *configs.Config
	//Logger
	MysqlDB      mysql.DB
	PostgresqlDB postgresql.DB
	RedisClient  redis.DB
}

func App(env configs.Env) *Application {
	app := &Application{}
	app.Config = ConfigLoad(env)
	//app.MysqlDB = NewMysqlDB(app.Config.Mysql)
	app.PostgresqlDB = NewPostgresqlDB(app.Config.Postgresql)
	app.RedisClient = NewRedisClient(app.Config.Redis)

	return app
}

func (a *Application) ClosePostgresqlConnection() {
	ClosePostgresqlDB(a.PostgresqlDB)
}

func (a *Application) CloseMysqlConnection() {
	CloseMysqlDB(a.MysqlDB)
}

func (a *Application) CloseRedisClientConnection() {
	CloseRedisClient(a.RedisClient)
}
