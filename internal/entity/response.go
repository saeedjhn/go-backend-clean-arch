package entity

type ErrorResponse struct {
	Success bool          `json:"Success"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors"`
	Meta    interface{}   `json:"meta,omitempty"`
}

func NewErrorResponse(
	message string,
	errors interface{},
	meta ...interface{},
) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Message: message,
		Errors:  []interface{}{errors},
		Meta:    meta,
	}
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

func NewSuccessResponse(
	message string,
	data interface{},
	meta ...interface{},
) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}
