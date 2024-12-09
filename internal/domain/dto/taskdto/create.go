package taskdto

type CreateRequest struct {
	UserID      uint64 `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Status      uint8  `json:"status"`
}

type CreateResponse struct {
	Data TaskInfo `json:"data"`
}
