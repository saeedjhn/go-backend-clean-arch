package bootstrap

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/configs"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/logger"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/cache/redis"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/mysql"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/persistance/db/pq"
)

type Application struct {
	Config      *configs.Config
	Logger      *logger.Logger
	MysqlDB     mysql.DB
	PostgresDB  pq.DB
	RedisClient redis.DB
}

func App(env configs.Env) *Application {
	app := &Application{}
	app.Config = ConfigLoad(env)
	app.Logger = newLogger(app.Config.Logger)
	app.MysqlDB = newMysqlConnection(app.Config.Mysql)
	app.PostgresDB = newPostgresConnection(app.Config.Postgres)
	app.RedisClient = newRedisClient(app.Config.Redis)

	return app
}

func (a *Application) ClosePostgresqlConnection() {
	closePostgresConnection(a.PostgresDB)
}

func (a *Application) CloseMysqlConnection() {
	closeMysqlConnection(a.MysqlDB)
}

func (a *Application) CloseRedisClientConnection() {
	closeRedisClient(a.RedisClient)
}
