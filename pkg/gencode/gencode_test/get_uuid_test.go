package gencode_test

import (
	"testing"

	"github.com/saeedjhn/go-domain-driven-design/pkg/gencode"
)

//go:generate go test -v -race -count=1 -run TestGenUUID

func TestGenUUID_GeneratesValidUUID_Success(t *testing.T) {
	uuid := gencode.GenUUID()
	if len(uuid) != 36 {
		t.Error("Expected UUID to be 36 characters long, got", len(uuid))
	}
}

func TestGenUUID_GeneratesUniqueUUIDs_Success(t *testing.T) {
	uuid1 := gencode.GenUUID()
	uuid2 := gencode.GenUUID()
	if uuid1 == uuid2 {
		t.Error("Expected UUIDs to be unique, got duplicates")
	}
}
