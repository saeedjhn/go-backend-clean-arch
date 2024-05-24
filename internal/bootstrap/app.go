package bootstrap

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/configs"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/logger"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/cache/redis"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/mysql"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/postgresql"
)

type Application struct {
	Config       *configs.Config
	Logger       *logger.Logger
	MysqlDB      mysql.DB
	PostgresqlDB postgresql.DB
	RedisClient  redis.DB
}

func App(env configs.Env) *Application {
	app := &Application{}
	app.Config = ConfigLoad(env)
	app.Logger = logger.New(app.Config.Logger)
	//app.MysqlDB = newMysqlConnection(app.Config.Mysql)
	app.PostgresqlDB = newPostgresqlConnection(app.Config.Postgresql)
	app.RedisClient = newRedisClient(app.Config.Redis)

	return app
}

func (a *Application) ClosePostgresqlConnection() {
	closePostgresqlConnection(a.PostgresqlDB)
}

func (a *Application) CloseMysqlConnection() {
	closeMysqlConnection(a.MysqlDB)
}

func (a *Application) CloseRedisClientConnection() {
	closeRedisClient(a.RedisClient)
}
