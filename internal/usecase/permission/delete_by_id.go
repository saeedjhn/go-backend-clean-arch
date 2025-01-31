package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/permission"
)

func (i *Interactor) DeleteByID(
	ctx context.Context,
	req permissiondto.DeleteByIDRequest,
) (permissiondto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
