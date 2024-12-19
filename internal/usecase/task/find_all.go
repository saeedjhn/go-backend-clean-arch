package task

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"
)

func (i *Interactor) FindAll(_ context.Context, _ task.FindAllRequest) (task.FindAllResponse, error) {
	panic("IMPLEMENT ME")
}
