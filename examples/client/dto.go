package main

type Request struct {
	ID uint64
}

type Response struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
