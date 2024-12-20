package task

type CreateRequest struct {
	UserID      uint64 `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Status      uint8  `json:"status"`
}

type CreateResponse struct {
	Data        Data              `json:"data"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
