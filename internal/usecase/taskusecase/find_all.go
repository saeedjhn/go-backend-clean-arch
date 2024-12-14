package taskusecase

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/taskdto"
)

func (i *Interactor) FindAll(_ context.Context, _ taskdto.FindAllRequest) (taskdto.FindAllResponse, error) {
	panic("IMPLEMENT ME")
}
