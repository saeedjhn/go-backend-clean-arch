package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/permission"
)

func (i *Interactor) DeleteAll(
	ctx context.Context,
	req permissiondto.DeleteAllRequest,
) (permissiondto.DeleteAllResponse, error) {
	panic("IMPLEMENT ME")
}
