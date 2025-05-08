package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/permission"
)

func (i Interactor) Create(
	_ context.Context,
	_ permissiondto.CreateRequest,
) (permissiondto.CreateResponse, error) {
	panic("IMPLEMENT ME")
}
