package sanitize_test

import (
	"github.com/saeedjhn/go-backend-clean-arch/pkg/sanitize"
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
	var err error
	s := sanitize.New().SetPolicy(sanitize.StrictPolicy)

	p := Post{
		Title:       "This is a <a>Title</a>",
		Description: "<b>Description</b>",
		Status:      Pending,
	}

	t.Log(p)

	res, err := s.Any(p)
	if err != nil {
		log.Println(err)
	}

	t.Log("result: ", res)

	if err = s.Struct(&p); err != nil {
		t.Fatal(err)
	}

	t.Log(p)
}
