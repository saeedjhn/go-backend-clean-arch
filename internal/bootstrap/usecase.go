package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	mysqluser "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	redisuser "github.com/saeedjhn/go-backend-clean-arch/internal/repository/redis/user"
	contract2 "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
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
	_ contract2.Logger,
	trc contract2.Tracer,
	cache Cache,
	db DB,
) *Usecase {
	var (
		userRepo    = mysqluser.New(trc, db.MySQL)
		userRdsRepo = redisuser.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)
	)

	var (
		authIntr = authusecase.New(config.Auth)
		userVld  = uservalidator.New(config.Application.EntropyPassword)
		userIntr = userusecase.New(config, trc, authIntr, userVld, userRdsRepo, userRepo)
	)

	return &Usecase{
		AuthIntr: authIntr,
		UserIntr: userIntr,
	}
}
