package role

import (
	"context"

	roledto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/role"
)

func (i Interactor) GetAll(_ context.Context, _ roledto.GetAllRequest) (roledto.GetAllResponse, error) {
	panic("IMPLEMENT ME")
}
