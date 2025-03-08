package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/roleresourcepermission"
)

func (i *Interactor) DeleteAll(
	_ context.Context,
	_ roleresourcepermissiondto.DeleteAllRequest,
) (roleresourcepermissiondto.DeleteAllResponse, error) {
	panic("IMPLEMENT ME")
}
