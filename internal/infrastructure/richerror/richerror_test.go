package richerror

import (
	"log"
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	messages := []string{"a", "b"}

	log.Println(strings.Join(messages, " "))
}
