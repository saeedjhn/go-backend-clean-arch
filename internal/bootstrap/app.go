package bootstrap

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/saeedjhn/go-backend-clean-arch/internal/buildinfo"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/cache/redis"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

// type Cache struct {
// 	InMem *inmemory.DB
// 	Redis *redis.DB
// }

// type DB struct {
// 	MySQL    *mysql.DB
// 	Postgres *pq.DB
// }

type Application struct {
	Config        *configs.Config
	BuildInfo     buildinfo.Info
	EventRegister types.EventRouter
	Logger        contract.Logger
	Trc           contract.Tracer
	Collector     contract.Collector
	Rabbitmq      contract.PublisherConsumer
	Redis         *redis.DB
	MySQL         *mysql.DB
	Usecase       *Usecase
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

	a.BuildInfo = NewBuildInfo()

	a.EventRegister = NewEventRegister()

	a.Logger = NewLogger(a.Config.Application, a.Config.Logger)

	if a.Trc, err = NewTracer(a.Config, a.BuildInfo); err != nil {
		return err
	}

	if a.Collector, err = NewCollector(a.Config, a.BuildInfo); err != nil {
		return err
	}

	if a.Rabbitmq, err = NewRabbitmq(a.Config.RabbitMQ, a.EventRegister); err != nil {
		return err
	}

	if a.MySQL, err = NewMysqlConnection(a.Config.Mysql); err != nil {
		return err
	}

	// if a.DB.Postgres, err = NewPostgresConnection(a.Config.Postgres); err != nil {
	// 	return err
	// }

	if a.Redis, err = NewRedisClient(a.Config.Redis); err != nil {
		return err
	}

	a.Usecase = NewUsecase(
		a.Config,
		a.Logger,
		a.Trc,
		a.MySQL,
	)

	return nil
}

func (a *Application) CloseMysqlConnection() error {
	return CloseMysqlConnection(a.MySQL)
}

func (a *Application) CloseRedisClientConnection() error {
	return CloseRedisClient(a.Redis)
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
