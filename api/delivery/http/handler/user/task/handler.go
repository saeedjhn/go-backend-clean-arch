package task

import (
	"context"

	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
	taskusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/task"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"
)

type Interactor interface {
	Create(ctx context.Context, req task.CreateRequest) (task.CreateResponse, error)
	GetAllByUserID(ctx context.Context, req task.GetAllByUserIDRequest) (task.GetByUserIDResponse, error)
}

type Handler struct {
	trc contract.Tracer
	// taskIntr Interactor
	authIntr *authusecase.Interactor
	taskIntr *taskusecase.Interactor
}

func New(
	trc contract.Tracer,
	// taskIntr Interactor,
	authIntr *authusecase.Interactor,
	taskIntr *taskusecase.Interactor,
) *Handler {
	return &Handler{
		trc:      trc,
		authIntr: authIntr,
		taskIntr: taskIntr,
	}
}
