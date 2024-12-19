package task

import (
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/http/handler/task"
)

type Validator struct {
}

var _ task.Validator = (*Validator)(nil)

func New() Validator {
	return Validator{}
}
