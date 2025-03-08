package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-domain-driven-design/internal/dto/resource"
)

func (i *Interactor) Create(_ context.Context, _ resourcedto.CreateRequest) (resourcedto.CreateResponse, error) {
	panic("IMPLEMENT ME")
}
