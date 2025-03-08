package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-domain-driven-design/internal/dto/resource"
)

func (i *Interactor) DeleteByID(
	_ context.Context,
	_ resourcedto.DeleteByIDRequest,
) (resourcedto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
