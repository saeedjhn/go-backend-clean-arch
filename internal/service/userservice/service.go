package userservice

import (
	"github.com/saeedjhn/go-backend-clean-arch/api/httpserver/handler/userhandler"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	userauthservicedto2 "github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	usertaskservicedto2 "github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type TaskGenerator interface {
	Create(dto usertaskservicedto2.CreateTaskRequest) (usertaskservicedto2.CreateTaskResponse, error)
	TasksUser(dto usertaskservicedto2.TasksUserRequest) (usertaskservicedto2.TasksUserResponse, error)
}

type AuthGenerator interface {
	CreateAccessToken(dto userauthservicedto2.CreateTokenRequest) (userauthservicedto2.CreateTokenResponse, error)
	CreateRefreshToken(dto userauthservicedto2.CreateTokenRequest) (userauthservicedto2.CreateTokenResponse, error)
	ExtractIDFromRefreshToken(dto userauthservicedto2.ExtractIDFromTokenRequest) (userauthservicedto2.ExtractIDFromTokenResponse, error)
}

type Repository interface {
	Create(u entity.User) (entity.User, error)
	IsMobileUnique(mobile string) (bool, error)
	GetByMobile(mobile string) (entity.User, error)
	GetByID(id uint) (entity.User, error)
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
