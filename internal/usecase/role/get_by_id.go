package role

import (
	"context"

	roledto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/role"
)

func (i Interactor) GetByID(_ context.Context, _ roledto.GetByIDRequest) (roledto.GetByIDResponse, error) {
	panic("IMPLEMENT ME")
}
