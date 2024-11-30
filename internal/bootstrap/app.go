package bootstrap

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract/tracercontract"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/logger"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/inmemory"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/redis"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/pq"
)

type Cache struct {
	InMem *inmemory.InMemory
	Redis *redis.Redis
}

type DB struct {
	MySQL    *mysql.Mysql
	Postgres *pq.Postgres
}

type Application struct {
	Config       *configs.Config
	ConfigOption configs.Option
	Trc          tracercontract.Tracer
	Logger       *logger.Logger
	Cache        Cache
	DB           DB
	Usecase      *Usecase
}

func App(configOption configs.Option) (*Application, error) {
	a := &Application{ConfigOption: configOption}

	if err := a.setup(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Application) setup() error {
	var err error

	if a.Config, err = ConfigLoad(a.ConfigOption); err != nil {
		return err
	}

	if a.DB.MySQL, err = NewMysqlConnection(a.Config.Mysql); err != nil {
		return err
	}

	if a.Trc, err = NewTracer(a.Config.Tracer); err != nil {
		return err
	}

	a.Logger = NewLogger(a.Config.Logger)

	if a.DB.Postgres, err = NewPostgresConnection(a.Config.Postgres); err != nil {
		return err
	}

	if a.Cache.Redis, err = NewRedisClient(a.Config.Redis); err != nil {
		return err
	}

	a.Cache.InMem = NewInMemory()

	a.Usecase = NewUsecase(
		a.Config,
		a.Logger,
		a.Cache,
		a.DB,
	)

	return nil
}

func (a *Application) CloseMysqlConnection() error {
	return CloseMysqlConnection(a.DB.MySQL)
}

func (a *Application) CloseRedisClientConnection() error {
	return CloseRedisClient(a.Cache.Redis)
}

func (a *Application) ShutdownTracer(ctx context.Context) error {
	return ShutdownTracer(ctx, a.Trc)
}

// func (a *Application) ClosePostgresqlConnection() error {
//	return ClosePostgresConnection(a.Postgres)
// }
