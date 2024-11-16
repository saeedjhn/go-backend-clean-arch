package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/userrespository/mysqluser"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/userrespository/redisuser"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authusecase"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/taskusecase"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/userusecase"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/logger"
)

type Usecase struct {
	AuthIntr *authusecase.Interactor
	UserIntr *userusecase.Interactor
	TaskIntr *taskusecase.Interactor
}

func NewUsecase(
	cfg *configs.Config,
	logger *logger.Logger,
	cache Cache,
	db DB,
) *Usecase {
	// Repositories
	taskRepo := mysqltask.New(db.MySQL)
	userRepo := mysqluser.New(db.MySQL)
	userRdsRepo := redisuser.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)

	// Dependencies

	// Usecase
	taskIntr := taskusecase.New(cfg, taskRepo)
	authIntr := authusecase.New(cfg.Auth)
	userIntr := userusecase.New(
		cfg,
		authIntr,
		taskIntr,
		userRdsRepo,
		userRepo,
	)

	return &Usecase{
		AuthIntr: authIntr,
		UserIntr: userIntr,
		TaskIntr: taskIntr,
	}
}
