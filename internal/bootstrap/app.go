package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/logger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
)

type Application struct {
	Config      *configs.Config
	Logger      *logger.Logger
	MysqlDB     mysql.DB
	PostgresDB  pq.DB
	RedisClient redis.DB
}

func App(env configs.Env) *Application {
	var app = &Application{}
	app.Config = ConfigLoad(env)
	app.Logger = NewLogger(app.Config.Logger)
	app.MysqlDB = NewMysqlConnection(app.Config.Mysql)
	//app.PostgresDB = NewPostgresConnection(app.Config.Postgres)
	app.RedisClient = NewRedisClient(app.Config.Redis)

	return app
}

//func (a *Application) ClosePostgresqlConnection() {
//	ClosePostgresConnection(a.PostgresDB)
//}

func (a *Application) CloseMysqlConnection() {
	CloseMysqlConnection(a.MysqlDB)
}

func (a *Application) CloseRedisClientConnection() {
	CloseRedisClient(a.RedisClient)
}
