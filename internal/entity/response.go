package entity

import "maps"

type ErrorResponse struct {
	Success bool                   `json:"Success"`
	Message string                 `json:"message"`
	Errors  []interface{}          `json:"errors"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

func NewErrorResponse(
	message string,
	errors interface{},
) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Message: message,
		Errors:  []interface{}{errors},
	}
}

func (e *ErrorResponse) WithMeta(meta map[string]interface{}) *ErrorResponse {
	maps.Copy(e.Meta, meta)

	return e
}

type SuccessResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

func NewSuccessResponse(
	message string,
	data interface{},
) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func (e *SuccessResponse) WithMeta(meta map[string]interface{}) *SuccessResponse {
	maps.Copy(e.Meta, meta)

	return e
}