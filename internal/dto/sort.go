package dto

type SortDirection string

const (
	AscSortDirection  = SortDirection("asc")
	DescSortDirection = SortDirection("desc")
)

type SortRequest struct {
	Field     string        `query:"sort_field"`
	Direction SortDirection `query:"sort_direction"`
}
