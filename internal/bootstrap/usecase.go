package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	usermysql "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
)

type Usecase struct {
	AuthIntr *authusecase.Interactor
	UserIntr *userusecase.Interactor
}

func NewUsecase(
	config *configs.Config,
	_ contract.Logger,
	trc contract.Tracer,
	_ Cache,
	db DB,
) *Usecase {
	var (
		userRepo = usermysql.New(trc, db.MySQL)
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
