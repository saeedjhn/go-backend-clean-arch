package richerror

import (
	"errors"
	"github.com/rotisserie/eris"
	. "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
)

type Op string

type RichError struct {
	op           Op
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
	stackTrace   map[string]interface{}
}

func (e RichError) Op() Op {
	return e.op
}

func (e RichError) WrappedError() error {
	return e.wrappedError
}

func (e RichError) Error() string {
	var err error
	if errors.As(e.wrappedError, &err) {
		return e.wrappedError.Error()
	}

	return ""
}

func (e RichError) Message() string {
	return e.message
}

func (e RichError) Kind() Kind {
	return e.kind
}

func (e RichError) Meta() map[string]interface{} {
	return e.meta
}

func (e RichError) StackTrace() map[string]interface{} {
	return e.stackTrace
}

func (e RichError) Get() map[string]interface{} {
	return map[string]interface{}{
		"op":          e.Op(),
		"error":       e.Error(),
		"message":     e.Message(),
		"kind":        e.Kind(),
		"meta":        e.Meta(),
		"stack_trace": e.StackTrace(),
	}
}

type RichErrorBuilder struct {
	RichError
}

func New(op Op) RichErrorBuilder {
	return RichErrorBuilder{RichError{op: op}}
}

func (r RichErrorBuilder) WithOp(op Op) RichErrorBuilder {
	r.op = op

	return r
}

func (r RichErrorBuilder) WithErr(err error) RichErrorBuilder {
	r.wrappedError = err

	return r
}

func (r RichErrorBuilder) WithMessage(message string) RichErrorBuilder {
	r.message = message

	return r
}

func (r RichErrorBuilder) WithKind(kind Kind) RichErrorBuilder {
	r.kind = kind

	return r
}

func (r RichErrorBuilder) WithMeta(meta map[string]interface{}) RichErrorBuilder {
	r.meta = meta

	return r
}

func (r RichErrorBuilder) WithStackTrace(message string) RichErrorBuilder {
	e := eris.ToJSON(eris.New(message), true)
	r.stackTrace = e

	return r
}

func (r RichErrorBuilder) Error() string {
	if r.message == "" && r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return r.message
}

func (r RichErrorBuilder) Op() Op {
	if r.op != "" {
		return r.op
	}

	var re RichErrorBuilder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return ""
	}

	return re.Op()
}

func (r RichErrorBuilder) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	var re RichErrorBuilder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichErrorBuilder) WrappedError() error {
	if r.wrappedError != nil {
		return r.wrappedError
	}

	var re RichErrorBuilder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return nil
	}

	return re.WrappedError()
}

func (r RichErrorBuilder) Message() string {
	if r.message != "" {
		return r.message
	}

	var re RichErrorBuilder
	ok := errors.As(r.wrappedError, &re)
	if ok {
		return re.Message()
	}

	if r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return ""
}

func (r RichErrorBuilder) Meta() map[string]interface{} {
	if len(r.meta) != 0 {
		return r.meta
	}

	var re RichErrorBuilder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return make(map[string]interface{})
	}

	return re.Meta()
}

func (r RichErrorBuilder) Build() RichError {
	return r.RichError
}

func Analysis(err error) (RichError, error) {
	var richError RichErrorBuilder
	switch {
	case errors.As(err, &richError):
		var re RichErrorBuilder
		errors.As(err, &re)

		return RichError{
			op:           re.Op(),
			wrappedError: re.WrappedError(),
			message:      re.Message(),
			kind:         re.Kind(),
			meta:         re.Meta(),
			stackTrace:   re.StackTrace(),
		}, nil

	default:
		return RichError{}, err
	}
}
