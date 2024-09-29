package userservice

import (
	"github.com/saeedjhn/go-backend-clean-arch/api/httpserver/handler/userhandler"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type TaskGenerator interface {
	Create(dto usertaskservicedto.CreateTaskRequest) (usertaskservicedto.CreateTaskResponse, error)
	TasksUser(dto usertaskservicedto.TasksUserRequest) (usertaskservicedto.TasksUserResponse, error)
}

type AuthGenerator interface {
	CreateAccessToken(dto userauthservicedto.CreateTokenRequest) (userauthservicedto.CreateTokenResponse, error)
	CreateRefreshToken(dto userauthservicedto.CreateTokenRequest) (userauthservicedto.CreateTokenResponse, error)
	ExtractIDFromRefreshToken(
		dto userauthservicedto.ExtractIDFromTokenRequest,
	) (userauthservicedto.ExtractIDFromTokenResponse, error)
}

type Repository interface {
	Create(u entity.User) (entity.User, error)
	IsMobileUnique(mobile string) (bool, error)
	GetByMobile(mobile string) (entity.User, error)
	GetByID(id uint64) (entity.User, error)
}

type UserInteractor struct {
	config         *configs.Config
	authInteractor AuthGenerator
	taskInteractor TaskGenerator
	repository     Repository
}

var _ userhandler.Interactor = (*UserInteractor)(nil)

func New(
	config *configs.Config,
	authInteractor AuthGenerator,
	taskInteractor TaskGenerator,
	repository Repository,
) *UserInteractor {
	return &UserInteractor{
		config:         config,
		authInteractor: authInteractor,
		taskInteractor: taskInteractor,
		repository:     repository,
	}
}
