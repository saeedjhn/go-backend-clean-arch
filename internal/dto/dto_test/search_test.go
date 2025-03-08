package dto_test

import (
	"testing"

	"github.com/saeedjhn/go-domain-driven-design/internal/dto"
	"github.com/stretchr/testify/assert"
)

//go:generate go test -v -race -count=1

func TestSearchRequest_AllowedFields_ValidFieldsSet(t *testing.T) {
	s := &dto.SearchRequest{Query: "test"}
	validFields := []string{"name", "email", "phone"}

	s.AllowedFields(validFields)

	assert.Equal(t, validFields, s.GetFields())
}

func TestSearchRequest_AllowedFields_EmptyQuery_NoFieldsSet(t *testing.T) {
	s := &dto.SearchRequest{Query: ""}

	s.AllowedFields([]string{"name", "email"})

	assert.Empty(t, s.GetFields())
}

func TestSearchRequest_GetSearch_QuerySanitized(t *testing.T) {
	s := &dto.SearchRequest{
		Query: " test@#123! ",
	}
	s.AllowedFields([]string{"name", "email"})

	expectedQuery := "test123"
	expectedResult := dto.QuerySearch{"name": expectedQuery, "email": expectedQuery}

	result := s.GetSearch()

	assert.Equal(t, expectedResult, *result)
}

func TestSearchRequest_GetSearch_EmptyQuery_ReturnsEmptyMap(t *testing.T) {
	s := &dto.SearchRequest{
		Query: "",
	}
	s.AllowedFields([]string{"name", "email"})

	result := s.GetSearch()

	assert.Empty(t, *result)
}

func TestSearchRequest_GetSearch_NoFields_ReturnsEmptyMap(t *testing.T) {
	s := &dto.SearchRequest{
		Query: "test123",
	}

	result := s.GetSearch()

	assert.Empty(t, *result)
}

func TestSearchRequest_GetSearch_SpecialCharactersRemoved(t *testing.T) {
	s := &dto.SearchRequest{
		Query: "test!@#$%^&*()123",
	}

	s.AllowedFields([]string{"name"})

	expectedQuery := "test123"
	expectedResult := dto.QuerySearch{"name": expectedQuery}

	result := s.GetSearch()

	assert.Equal(t, expectedResult, *result)
}

func TestSearchRequest_GetSearch_MultipleFields_AllMapped(t *testing.T) {
	s := &dto.SearchRequest{
		Query: "searchText",
	}
	s.AllowedFields([]string{"name", "email", "phone"})

	expectedResult := dto.QuerySearch{
		"name":  "searchText",
		"email": "searchText",
		"phone": "searchText",
	}

	result := s.GetSearch()

	assert.Equal(t, expectedResult, *result)
}
