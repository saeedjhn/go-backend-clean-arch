package gencode_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/gencode"
)

//go:generate go test -v -race -count=1 ./...

func TestGenCode_ValidLength_ReturnsStringWithCorrectLength(t *testing.T) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 10

	code, err := gencode.GenCode(length, chars)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(code) != length {
		t.Errorf("expected length %d, got %d", length, len(code))
	}
}

func TestGenCode_EmptyChars_ReturnsError(t *testing.T) {
	_, err := gencode.GenCode(10, "")
	if err == nil {
		t.Errorf("expected error for empty chars, but got nil")
	}
}

func TestGenCode_ZeroLength_ReturnsError(t *testing.T) {
	chars := "abcdefghijklmnopqrstuvwxyz"
	_, err := gencode.GenCode(0, chars)
	if err == nil {
		t.Errorf("expected error for zero length, but got nil")
	}
}

func TestGenCode_NegativeLength_ReturnsError(t *testing.T) {
	chars := "abcdefghijklmnopqrstuvwxyz"
	_, err := gencode.GenCode(-5, chars)
	if err == nil {
		t.Errorf("expected error for negative length, but got nil")
	}
}
