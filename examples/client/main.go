package main

import (
	"context"
	"log"
)

type Request struct {
}

type Response struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//type PresenceClient interface {
//	GetPresence(ctx context.Context, request param.GetPresenceRequest) (param.GetPresenceResponse, error)
//}

type Client interface {
	Get(ctx context.Context, req Request) (Response, error)
}

func main() {
	addr := "https://jsonplaceholder.typicode.com/posts/{postId}"

	client := NewHTTPClient(addr)
	client.WithPath("postId", "1")
	resp, err := client.Get(context.Background(), Request{})
	if err != nil {
		log.Fatalf("Error client: %v", err)
	}

	log.Printf("%#v", resp)
}
