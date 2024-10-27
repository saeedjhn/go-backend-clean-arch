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
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/taskservice"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/userservice"
)

type Provider struct {
	AuthSvc *authservice.AuthInteractor
	UserSvc *userservice.UserInteractor
	TaskSvc *taskservice.TaskInteractor
}

func NewProvider(
	cfg *configs.Config,
	logger *logger.Logger,
	rds redis.DB,
	mySQLDB mysql.DB,
	pqDB pq.DB,
) *Provider {
	// Repositories
	taskRepo := mysqltask.New(mySQLDB)
	userRepo := mysqluser.New(mySQLDB)

	// Dependencies
	token := token.New()

	// Provider
	// Provider-oriented - inject service to another service
	taskSvc := taskservice.New(cfg, taskRepo)
	authSvc := authservice.New(cfg.Auth, token)
	userSvc := userservice.New(cfg, authSvc, taskSvc, userRepo)

	return &Provider{
		AuthSvc: authSvc,
		UserSvc: userSvc,
		TaskSvc: taskSvc,
	}
}
