package contract

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/examples/client/dto"
)

type Client interface {
	FetchByID(ctx context.Context, req dto.Request) (dto.Response, error)
}
