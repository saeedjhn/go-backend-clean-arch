package gencode_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/gencode"
)

//go:generate go test -v -race -count=1 -bench=. -benchmem -run BenchmarkGenCode

func BenchmarkGenCode_SmallLength(b *testing.B) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 10

	for i := 0; i < b.N; i++ {
		_, _ = gencode.GenCode(length, chars)
	}
}

func BenchmarkGenCode_MediumLength(b *testing.B) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 100

	for i := 0; i < b.N; i++ {
		_, _ = gencode.GenCode(length, chars)
	}
}

func BenchmarkGenCode_LargeLength(b *testing.B) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 1000

	for i := 0; i < b.N; i++ {
		_, _ = gencode.GenCode(length, chars)
	}
}
