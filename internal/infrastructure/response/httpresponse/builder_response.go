package httpresponse

import "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/response/httpresponse/internal"

type Builder struct {
	entity internal.HTTPResponse
}

func New() Builder {
	e := internal.HTTPResponse{}

	return Builder{entity: e}
}

func (h Builder) WithStatus(status bool) Builder {
	h.entity.Status = status

	return h
}

func (h Builder) WithStatusCode(statusCode int) Builder {
	h.entity.StatusCode = statusCode

	return h
}

func (h Builder) WithRequestID(requestID string) Builder {
	h.entity.RequestID = requestID

	return h
}

func (h Builder) WithPath(path string) Builder {
	h.entity.Path = path

	return h
}

func (h Builder) WithExecutionDuration(execDuration string) Builder {
	h.entity.ExecutionDuration = execDuration

	return h
}

func (h Builder) WithMessage(message string) Builder {
	h.entity.Message = message

	return h
}

func (h Builder) WithMeta(meta interface{}) Builder {
	h.entity.Meta = meta

	return h
}

func (h Builder) WithError(err interface{}) Builder {
	h.entity.Meta = map[string]interface{}{"error": err}

	return h
}

func (h Builder) Build() internal.HTTPResponse {
	return h.entity
}
