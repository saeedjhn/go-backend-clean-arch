package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/admin"
)

func (i *Interactor) RefreshToken(
	_ context.Context,
	_ admindto.RefreshTokenRequest,
) (admindto.RefreshTokenResponse, error) {
	panic("IMPLEMENT_ME")
}
