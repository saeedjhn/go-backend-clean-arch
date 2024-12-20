package task

type FindAllByUserIDRequest struct {
	UserID uint64 `param:"id" json:"user_id"`
}

type FindAllByUserIDResponse struct {
	Data        []Data            `json:"data"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
