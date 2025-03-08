package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/permission"
)

func (i *Interactor) Update(
	_ context.Context,
	_ permissiondto.UpdateRequest,
) (permissiondto.UpdateResponse, error) {
	panic("IMPLEMENT ME")
}
