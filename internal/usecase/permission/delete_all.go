package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/permission"
)

func (i *Interactor) DeleteAll(
	_ context.Context,
	_ permissiondto.DeleteAllRequest,
) (permissiondto.DeleteAllResponse, error) {
	panic("IMPLEMENT ME")
}
