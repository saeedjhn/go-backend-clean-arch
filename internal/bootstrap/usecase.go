package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	outboxevent "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/outbox_event"
	usermysql "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
	outboxusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/outbox"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

type Usecase struct {
	AuthIntr authusecase.Interactor
	UserIntr userusecase.Interactor
}

func NewUsecase(
	config *configs.Config,
	_ contract.Logger,
	trc contract.Tracer,
	mysqlDB *mysql.DB,
) *Usecase {
	var (
		userRepo   = usermysql.New(trc, mysqlDB)
		outboxRepo = outboxevent.New(mysqlDB)
		// userRdsRepo = userredis.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)
	)

	var (
		authIntr   = authusecase.New(config.Auth)
		outboxIntr = outboxusecase.New(config, trc, outboxRepo)
		userVld    = uservalidator.New(config.Application.EntropyPassword)
		userIntr   = userusecase.New(config, trc, authIntr, outboxIntr, userVld, userRepo)
	)

	return &Usecase{
		AuthIntr: authIntr,
		UserIntr: userIntr,
	}
}
