package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/permission"
)

func (i Interactor) GetByID(
	_ context.Context,
	_ permissiondto.GetByIDRequest,
) (permissiondto.GetByIDResponse, error) {
	panic("IMPLEMENT ME")
}
