package task

import (
	"context"

	"github.com/saeedjhn/go-domain-driven-design/internal/dto/task"
)

func (i *Interactor) GetAll(_ context.Context, _ task.FindAllRequest) (task.FindAllResponse, error) {
	panic("IMPLEMENT ME")
}
