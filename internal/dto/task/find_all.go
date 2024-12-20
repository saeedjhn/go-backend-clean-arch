package task

type FindAllRequest struct {
}

type FindAllResponse struct {
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
