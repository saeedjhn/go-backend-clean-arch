package main

import "context"

type GRPCClient struct {
	addr string
}

func NewGRPCClient(addr string) *GRPCClient {
	return &GRPCClient{addr: addr}
}

func (c GRPCClient) Get(ctx context.Context, req Request) (Response, error) {
	//TODO implement me
	panic("implement me")
}
