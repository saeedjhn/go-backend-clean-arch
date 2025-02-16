package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/roleresourcepermission"
)

func (i *Interactor) Create(
	_ context.Context,
	_ roleresourcepermissiondto.CreateRequest,
) (roleresourcepermissiondto.CreateResponse, error) {
	panic("IMPLEMENT ME")
}
