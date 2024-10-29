package bootstrap

import (
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/logger"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/mysql"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/persistance/db/pq"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/token"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/taskrepository/mysqltask"
	"github.com/saeedjhn/go-backend-clean-arch/internal/repository/userrespository/mysqluser"
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
	cfg *configs.Config,
	logger *logger.Logger,
	rds redis.DB,
	mySQLDB mysql.DB,
	pqDB pq.DB,
) *Usecase {
	// Repositories
	taskRepo := mysqltask.New(mySQLDB)
	userRepo := mysqluser.New(mySQLDB)

	// Dependencies
	token := token.New()

	// Usecase
	taskIntr := taskusecase.New(cfg, taskRepo)
	authIntr := authusecase.New(cfg.Auth, token)
	userIntr := userusecase.New(cfg, authIntr, taskIntr, userRepo)

	return &Usecase{
		AuthIntr: authIntr,
		UserIntr: userIntr,
		TaskIntr: taskIntr,
	}
}
