package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-domain-driven-design/internal/dto/resource"
)

func (i *Interactor) GetAll(_ context.Context, _ resourcedto.GetAllRequest) (resourcedto.GetAllResponse, error) {
	panic("IMPLEMENT ME")
}
