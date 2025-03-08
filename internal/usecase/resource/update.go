package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-domain-driven-design/internal/dto/resource"
)

func (i *Interactor) Update(_ context.Context, _ resourcedto.UpdateRequest) (resourcedto.UpdateResponse, error) {
	panic("IMPLEMENT ME")
}
