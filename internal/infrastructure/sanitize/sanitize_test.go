package sanitize

import (
	"log"
	"testing"
)

type Status string

const (
	Pending Status = "Pending"
)

type Post struct {
	Title       string
	Description string
	Status      Status
}

func TestSanitize(t *testing.T) {
	s := New().SetPolicy(StrictPolicy)

	p := Post{
		Title:       "This is a <a>Title</a>",
		Description: "<b>Description</b>",
		Status:      Pending,
	}

	log.Println(p)

	any, err := s.Any(p)
	if err != nil {
		log.Println(err)
	}
	log.Println("any: ", any)

	s.Struct(&p)

	log.Println(p)
}
