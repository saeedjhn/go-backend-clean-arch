package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/permission"
)

func (i *Interactor) DeleteByID(
	_ context.Context,
	_ permissiondto.DeleteByIDRequest,
) (permissiondto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
