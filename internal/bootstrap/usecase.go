package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	task2 "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/task"
	user2 "github.com/saeedjhn/go-backend-clean-arch/internal/repository/mysql/user"
	user3 "github.com/saeedjhn/go-backend-clean-arch/internal/repository/redis/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/auth"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/task"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
)

type Usecase struct {
	AuthIntr *auth.Interactor
	UserIntr *user.Interactor
	TaskIntr *task.Interactor
}

func NewUsecase(
	config *configs.Config,
	_ contract.Logger,
	trc contract.Tracer,
	cache Cache,
	db DB,
) *Usecase {
	// Repositories
	taskRepo := task2.New(db.MySQL)
	userRepo := user2.New(trc, db.MySQL)
	userRdsRepo := user3.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)

	// Dependencies

	// Usecase
	taskIntr := task.New(config, taskRepo)
	authIntr := auth.New(config.Auth)
	userIntr := user.New(
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
