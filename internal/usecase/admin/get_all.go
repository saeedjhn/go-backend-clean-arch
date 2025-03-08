package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-domain-driven-design/internal/dto/admin"
)

func (i *Interactor) GetAll(
	_ context.Context,
	_ admindto.GetAllRequest,
) (admindto.GetAllResponse, error) {
	panic("IMPLEMENT_ME")
}
