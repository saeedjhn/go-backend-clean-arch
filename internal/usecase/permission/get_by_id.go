package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/permission"
)

func (i *Interactor) GetByID(
	_ context.Context,
	_ permissiondto.GetByIDRequest,
) (permissiondto.GetByIDResponse, error) {
	panic("IMPLEMENT ME")
}
