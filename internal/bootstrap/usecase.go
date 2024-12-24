package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	mysqltask "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/task"
	mysqluser "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	redisuser "github.com/saeedjhn/go-backend-clean-arch/internal/repository/redis/user"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/auth"
	taskusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/task"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	uservalidator "github.com/saeedjhn/go-backend-clean-arch/internal/validator/user"
)

type Usecase struct {
	AuthIntr *authusecase.Interactor
	UserIntr *userusecase.Interactor
	TaskIntr *taskusecase.Interactor
}

func NewUsecase(
	config *configs.Config,
	_ contract.Logger,
	trc contract.Tracer,
	cache Cache,
	db DB,
) *Usecase {
	var (
		taskRepo    = mysqltask.New(db.MySQL)
		userRepo    = mysqluser.New(trc, db.MySQL)
		userRdsRepo = redisuser.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)
	)

	var (
		taskIntr = taskusecase.New(config, taskRepo)
		authIntr = authusecase.New(config.Auth)
		userVld  = uservalidator.New(config.Application.EntropyPassword)
		userIntr = userusecase.New(config, trc, userVld, authIntr, userRdsRepo, userRepo)
	)

	return &Usecase{
		AuthIntr: authIntr,
		UserIntr: userIntr,
		TaskIntr: taskIntr,
	}
}
