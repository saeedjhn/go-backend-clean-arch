package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/permission"
)

func (i *Interactor) DeleteByID(
	_ context.Context,
	_ permissiondto.DeleteByIDRequest,
) (permissiondto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
