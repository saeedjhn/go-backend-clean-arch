package userusecase

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

type TaskInteractor interface {
	Create(ctx context.Context, dto usertaskservicedto.CreateTaskRequest) (usertaskservicedto.CreateTaskResponse, error)
	TasksUser(ctx context.Context, dto usertaskservicedto.TasksUserRequest) (usertaskservicedto.TasksUserResponse, error)
}

type AuthInteractor interface {
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

type Interactor struct {
	config     *configs.Config
	authIntr   AuthInteractor
	taskIntr   TaskInteractor
	repository Repository
}

// var _ userhandler.Interactor = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(
	config *configs.Config,
	authIntr AuthInteractor,
	taskIntr TaskInteractor,
	repository Repository,
) *Interactor {
	return &Interactor{
		config:     config,
		authIntr:   authIntr,
		taskIntr:   taskIntr,
		repository: repository,
	}
}
