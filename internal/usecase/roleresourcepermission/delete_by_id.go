package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/roleresourcepermission"
)

func (i *Interactor) DeleteByID(
	_ context.Context,
	_ roleresourcepermissiondto.DeleteByIDRequest,
) (roleresourcepermissiondto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
