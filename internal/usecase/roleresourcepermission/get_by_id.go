package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/roleresourcepermission"
)

func (i *Interactor) GetByID(
	_ context.Context,
	_ roleresourcepermissiondto.GetByIDRequest,
) (roleresourcepermissiondto.GetByIDResponse, error) {
	panic("IMPLEMENT ME")
}
