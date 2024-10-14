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
	EnvMode     configs.Env
	Logger      *logger.Logger
	MySQLDB     mysql.DB
	PostgresDB  pq.DB
	RedisClient redis.DB
}

func App(env configs.Env) (*Application, error) {
	a := &Application{EnvMode: env}

	if err := a.setup(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Application) setup() error {
	var err error

	if a.Config, err = ConfigLoad(a.EnvMode); err != nil {
		return err
	}

	if a.MySQLDB, err = NewMysqlConnection(a.Config.Mysql); err != nil {
		return err
	}

	a.Logger = NewLogger(a.Config.Logger)

	if a.PostgresDB, err = NewPostgresConnection(a.Config.Postgres); err != nil {
		return err
	}

	if a.RedisClient, err = NewRedisClient(a.Config.Redis); err != nil {
		return err
	}

	return nil
}

func (a *Application) CloseMysqlConnection() error {
	return CloseMysqlConnection(a.MySQLDB)
}

func (a *Application) CloseRedisClientConnection() error {
	return CloseRedisClient(a.RedisClient)
}

// func (a *Application) ClosePostgresqlConnection() error {
//	return ClosePostgresConnection(a.Postgres)
// }
