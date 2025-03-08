package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/permission"
)

func (i *Interactor) GetAll(
	_ context.Context,
	_ permissiondto.GetAllRequest,
) (permissiondto.GetAllResponse, error) {
	panic("IMPLEMENT ME")
}
