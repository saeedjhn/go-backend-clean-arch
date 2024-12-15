package main

import (
	"context"
)

type Client interface {
	FetchByID(ctx context.Context, req Request) (Response, error)
}
