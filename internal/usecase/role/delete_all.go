package role

import (
	"context"

	roledto "github.com/saeedjhn/go-domain-driven-design/internal/dto/role"
)

func (i *Interactor) DeleteAll(_ context.Context, _ roledto.DeleteAllRequest) (roledto.DeleteAllResponse, error) {
	panic("IMPLEMENT ME")
}
