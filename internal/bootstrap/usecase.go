package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/userrespository/mysqluser"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/userrespository/redisuser"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authusecase"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/taskusecase"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/userusecase"
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
	// Repositories
	taskRepo := mysqltask.New(db.MySQL)
	userRepo := mysqluser.New(trc, db.MySQL)
	userRdsRepo := redisuser.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)

	// Dependencies

	// Usecase
	taskIntr := taskusecase.New(config, taskRepo)
	authIntr := authusecase.New(config.Auth)
	userIntr := userusecase.New(
		config,
		trc,
		authIntr,
		userRdsRepo,
		userRepo,
	)

	return &Usecase{
		AuthIntr: authIntr,
		UserIntr: userIntr,
		TaskIntr: taskIntr,
	}
}
