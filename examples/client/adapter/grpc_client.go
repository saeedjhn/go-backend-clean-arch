package adapter

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/examples/client/dto"
)

type GRPCClient struct {
	addr string
}

func NewGRPCClient(addr string) *GRPCClient {
	return &GRPCClient{addr: addr}
}

func (c GRPCClient) FetchByID(ctx context.Context, req dto.Request) (dto.Response, error) {
	// TODO implement me
	panic("implement me")
}
