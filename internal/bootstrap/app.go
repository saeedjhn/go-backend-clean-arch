package bootstrap

import (
	"context"

	contract2 "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/redis"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/inmemory"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/pq"
)

type Cache struct {
	InMem *inmemory.DB
	Redis *redis.DB
}

type DB struct {
	MySQL    *mysql.DB
	Postgres *pq.DB
}

type Application struct {
	Config    *configs.Config
	Logger    contract2.Logger
	Trc       contract2.Tracer
	Collector contract2.Collector
	Cache     Cache
	DB        DB
	Usecase   *Usecase
}

func App(config *configs.Config) (*Application, error) {
	a := &Application{Config: config}

	if err := a.setup(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Application) setup() error {
	var err error

	a.Logger = NewLogger(a.Config.Application, a.Config.Logger)

	if a.Trc, err = NewTracer(
		a.Config.Tracer,
		a.Config.Application,
		a.Config.HTTPServer,
	); err != nil {
		return err
	}

	if a.Collector, err = NewCollector(
		a.Config.Collector,
		a.Config.Application,
		a.Config.HTTPServer,
	); err != nil {
		return err
	}

	if a.DB.MySQL, err = NewMysqlConnection(a.Config.Mysql); err != nil {
		return err
	}

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
		a.Trc,
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

func (a *Application) ShutdownCollector(ctx context.Context) error {
	return ShutdownCollector(ctx, a.Collector)
}

// func (a *Application) ClosePostgresqlConnection() error {
//	return ClosePostgresConnection(a.DB)
// }
