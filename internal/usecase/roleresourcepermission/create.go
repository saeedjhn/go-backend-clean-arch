package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/roleresourcepermission"
)

func (i *Interactor) Create(
	_ context.Context,
	_ roleresourcepermissiondto.CreateRequest,
) (roleresourcepermissiondto.CreateResponse, error) {
	panic("IMPLEMENT ME")
}
