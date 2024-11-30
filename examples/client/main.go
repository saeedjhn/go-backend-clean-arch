package main

import (
	"context"
	"fmt"
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/examples/client/adapter"
	"github.com/saeedjhn/go-backend-clean-arch/examples/client/contract"
	"github.com/saeedjhn/go-backend-clean-arch/examples/client/dto"
)

// type PresenceClient interface {
//	GetPresence(ctx context.Context, request param.GetPresenceRequest) (param.GetPresenceResponse, error)
// }

type Service struct {
	client contract.Client
}

func New(client contract.Client) *Service {
	return &Service{client: client}
}

func (s Service) GetByID(ctx context.Context, req dto.Request) (dto.Response, error) {
	resp, err := s.client.FetchByID(ctx, req)
	if err != nil {
		return dto.Response{}, fmt.Errorf("error from client: %w", err)
	}

	return resp, nil
}

func main() {
	addr := "https://jsonplaceholder.typicode.com/posts/{postId}"

	client := adapter.NewHTTPClient(addr)
	// client.WithPath("postId", "1")

	req := dto.Request{ID: 1}
	svc := New(client)
	resp, err := svc.GetByID(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", resp)

	// resp, err := client.Get(context.Background(), Request{})
	// if err != nil {
	//	log.Fatalf("Error client: %v", err)
	// }
	// log.Printf("%#v", resp)
}
