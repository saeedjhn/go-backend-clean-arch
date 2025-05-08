package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	usermysql "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

type Usecase struct {
	AuthIntr *authusecase.Interactor
	UserIntr *userusecase.Interactor
}

func NewUsecase(
	config *configs.Config,
	_ contract.Logger,
	trc contract.Tracer,
	mysqlDB *mysql.DB,
) *Usecase {
	var (
		userRepo = usermysql.New(trc, mysqlDB)
		// userRdsRepo = userredis.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)
	)

	var (
		authIntr = authusecase.New(config.Auth)
		userVld  = uservalidator.New(config.Application.EntropyPassword)
		userIntr = userusecase.New(config, trc, authIntr, userVld, userRepo)
	)

	return &Usecase{
		AuthIntr: authIntr,
		UserIntr: userIntr,
	}
}
