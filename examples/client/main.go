package main

import (
	"context"
	"fmt"
	"log"
)

// type PresenceClient interface {
//	GetPresence(ctx context.Context, request param.GetPresenceRequest) (param.GetPresenceResponse, error)
// }

type Service struct {
	client Client
}

func New(client Client) *Service {
	return &Service{client: client}
}

func (s Service) GetByID(ctx context.Context, req Request) (Response, error) {
	resp, err := s.client.FetchByID(ctx, req)
	if err != nil {
		return Response{}, fmt.Errorf("error from client: %w", err)
	}

	return resp, nil
}

func main() {
	addr := "https://jsonplaceholder.typicode.com/posts/{postId}"

	client := NewHTTPAdaptor(addr)
	// client.WithPath("postId", "1")

	req := Request{ID: 1}
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
