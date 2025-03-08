package role

import (
	"context"

	roledto "github.com/saeedjhn/go-domain-driven-design/internal/dto/role"
)

func (i *Interactor) DeleteByID(_ context.Context, _ roledto.DeleteByIDRequest) (roledto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
