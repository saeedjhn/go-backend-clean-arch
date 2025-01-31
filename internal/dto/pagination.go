package dto

const (
	defaultPageNumber = 1
	defaultPageSize   = 10
)

type PaginationRequest struct {
	PageSize   uint `query:"page_size" example:"10"`
	PageNumber uint `query:"page_number" example:"1"`
}

type PaginationResponse struct {
	PageSize   uint `json:"page_size" example:"10"`
	PageNumber uint `json:"page_number" example:"1"`
	Total      uint `json:"total" example:"100"`
}

func (p *PaginationRequest) GetPageNumber() uint {
	if p.PageNumber <= 0 {
		p.PageNumber = defaultPageNumber
	}

	return p.PageNumber
}

func (p *PaginationRequest) GetOffset() uint {
	return (p.GetPageNumber() - 1) * p.GetPageSize()
}

func (p *PaginationRequest) GetPageSize() uint {
	validPageSizes := []uint{10, 25, 50, 100}
	for _, size := range validPageSizes {
		if p.PageSize == size {
			return p.PageSize
		}
	}
	p.PageSize = defaultPageSize

	return p.PageSize
}
