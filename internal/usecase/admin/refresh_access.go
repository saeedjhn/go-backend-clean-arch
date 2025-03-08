package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-domain-driven-design/internal/dto/admin"
)

func (i *Interactor) RefreshToken(
	_ context.Context,
	_ admindto.RefreshTokenRequest,
) (admindto.RefreshTokenResponse, error) {
	panic("IMPLEMENT_ME")
}
