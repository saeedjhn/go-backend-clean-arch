package task

type GetAllByUserIDRequest struct {
	UserID uint64 `param:"id" json:"user_id"`
}

type GetByUserIDResponse struct {
	Tasks       []TaskInfo        `json:"tasks"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
