package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-domain-driven-design/internal/dto/resource"
)

func (i *Interactor) DeleteAll(
	_ context.Context,
	_ resourcedto.DeleteAllRequest,
) (resourcedto.DeleteAllResponse, error) {
	panic("IMPLEMENT ME")
}
