package gencode_test

import (
	"testing"

	"github.com/saeedjhn/go-domain-driven-design/pkg/gencode"
)

//go:generate go test -v -race -count=1 -bench=. -benchmem -run BenchmarkGenCode

func BenchmarkGenCode_SmallLength(b *testing.B) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 10

	for range b.N {
		_, _ = gencode.GenCode(length, chars)
	}
}

func BenchmarkGenCode_MediumLength(b *testing.B) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 100

	for range b.N {
		_, _ = gencode.GenCode(length, chars)
	}
}

func BenchmarkGenCode_LargeLength(b *testing.B) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 1000

	for range b.N {
		_, _ = gencode.GenCode(length, chars)
	}
}
