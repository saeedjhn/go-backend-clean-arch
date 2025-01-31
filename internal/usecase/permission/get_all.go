package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/permission"
)

func (i *Interactor) GetAll(ctx context.Context, req permissiondto.GetAllRequest) (permissiondto.GetAllResponse, error) {
	panic("IMPLEMENT ME")
}
