package dto

import (
	"regexp"
	"strings"
)

type QuerySearch map[string]interface{}

type SearchRequest struct {
	fields []string
	Query  string `query:"search" example:"09123456789"`
}

func (s *SearchRequest) AllowedFields(fields []string) *SearchRequest {
	if s.Query == "" {
		return s
	}

	s.fields = fields

	return s
}

func (s *SearchRequest) GetFields() []string {
	return s.fields
}

func (s *SearchRequest) GetSearch() *QuerySearch {
	searchParams := QuerySearch{}

	if s.Query != "" {
		s.Query = strings.TrimSpace(s.Query)
		re := regexp.MustCompile(`[\p{P}\p{S}]+`)
		s.Query = re.ReplaceAllString(s.Query, "")
		s.Query = strings.ReplaceAll(s.Query, " ", "")

		for _, val := range s.fields {
			searchParams[val] = s.Query
		}
	}

	return &searchParams
}
