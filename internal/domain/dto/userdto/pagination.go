package userdto

type Pagination struct {
	CurrentPage int32 `json:"current_page"`
	TotalPages  int32 `json:"total_pages"`
	PerPage     int32 `json:"per_page"`
	TotalItems  int32 `json:"total_items"`
}
