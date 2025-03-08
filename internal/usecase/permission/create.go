package permission

import (
	"context"

	permissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/permission"
)

func (i *Interactor) Create(
	_ context.Context,
	_ permissiondto.CreateRequest,
) (permissiondto.CreateResponse, error) {
	panic("IMPLEMENT ME")
}
