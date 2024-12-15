package main

import (
	"context"
)

type GRPCAdaptor struct {
	addr string
}

func NewGRPCAdaptor(addr string) *GRPCAdaptor {
	return &GRPCAdaptor{addr: addr}
}

func (c GRPCAdaptor) FetchByID(ctx context.Context, req Request) (Response, error) {
	// TODO implement me
	panic("implement me")
}
