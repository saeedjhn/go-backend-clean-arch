package richerror

import (
	"encoding/json"
	"errors"
)

// Op represents the operation where the error occurred.
type Op string

// RichError is the main implementation of the RichErrorInterface.
type RichError struct {
	op           Op
	kind         Kind
	message      string
	wrappedError error
	meta         map[string]interface{}
}

// New creates a new RichError.
func New(op Op) *RichError {
	return &RichError{
		op:   op,
		meta: make(map[string]interface{}),
	}
}

// WithMessage sets the error message.
func (e *RichError) WithMessage(message string) *RichError {
	e.message = message
	return e
}

// WithKind sets the error kind.
func (e *RichError) WithKind(kind Kind) *RichError {
	e.kind = kind
	return e
}

// WithErr wraps another error.
func (e *RichError) WithErr(err error) *RichError {
	e.wrappedError = err
	return e
}

// WithMeta adds metadata to the error.
func (e *RichError) WithMeta(meta map[string]interface{}) *RichError {
	e.meta = meta
	return e
}

// WithTraceID adds a trace ID to the error metadata.
// func (e *RichError) WithTraceID(traceID string) *RichError {
// 	e.WithMeta("trace_id", traceID)
// 	return e
// }

// Op returns the operation where the error occurred.
// It tries to fetch the operation from wrapped errors if not set.
func (e *RichError) Op() Op {
	if e.op != "" {
		return e.op
	}

	var wrapped *RichError
	if errors.As(e.wrappedError, &wrapped) {
		return wrapped.Op() // Recursive call
	}

	return ""
}

// Kind returns the kind of the error.
// It checks wrapped errors recursively if the kind is not set.
func (e *RichError) Kind() Kind {
	if e.kind != KindUnknown {
		return e.kind
	}

	var wrapped *RichError
	if errors.As(e.wrappedError, &wrapped) {
		return wrapped.Kind() // Recursive call
	}

	return KindUnknown
}

// Message returns the error message.
// It falls back to wrapped errors recursively if no message is set.
func (e *RichError) Message() string {
	if e.message != "" {
		return e.message
	}

	var wrapped *RichError
	if errors.As(e.wrappedError, &wrapped) {
		return wrapped.Message() // Recursive call
	}

	return ""
}

// Meta aggregates metadata from the current and wrapped errors.
func (e *RichError) Meta() map[string]interface{} {
	meta := make(map[string]interface{})

	var wrapped *RichError
	if errors.As(e.wrappedError, &wrapped) {
		for k, v := range wrapped.Meta() { // Recursive call
			if _, exists := meta[k]; !exists {
				meta[k] = v
			}
		}
	}

	for k, v := range e.meta {
		meta[k] = v
	}

	return meta
}

// Error implements the error interface.
func (e *RichError) Error() string {
	if e.wrappedError != nil {
		return e.wrappedError.Error()
	}

	return e.message
}

// WrappedError returns the wrapped error.
func (e *RichError) WrappedError() error {
	return e.wrappedError
}

// ToJSON serializes the error to JSON.
func (e *RichError) ToJSON() (string, error) {
	data := map[string]interface{}{
		"op":      e.Op(),
		"kind":    e.Kind(),
		"message": e.Message(),
		"meta":    e.Meta(),
	}

	if e.wrappedError != nil {
		data["wrapped_error"] = e.wrappedError.Error()
	}

	jsonData, err := json.Marshal(data)

	return string(jsonData), err
}

// Analysis extracts RichError from a generic error.
func Analysis(err error) RichError {
	var richErr *RichError

	if errors.As(err, &richErr) {
		return RichError{
			op:           richErr.Op(),
			kind:         richErr.Kind(),
			message:      richErr.Message(),
			wrappedError: richErr.WrappedError(),
			meta:         richErr.Meta(),
		}
	}

	return RichError{}
}
