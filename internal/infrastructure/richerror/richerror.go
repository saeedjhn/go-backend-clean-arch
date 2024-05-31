package richerror

import (
	"errors"
	. "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
)

type Op string

type RichErr struct {
	op           Op
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
}

func (e RichErr) Op() Op {
	return e.op
}

func (e RichErr) WrappedError() error {
	return e.wrappedError
}

func (e RichErr) Message() string {
	return e.message
}

func (e RichErr) Kind() Kind {
	return e.kind
}

func (e RichErr) Meta() map[string]interface{} {
	return e.meta
}

func (e RichErr) Get() map[string]interface{} {
	return map[string]interface{}{
		"op":      e.Op(),
		"error":   e.WrappedError(),
		"message": e.Message(),
		"kind":    e.Kind(),
		"meta":    e.Meta(),
	}
}

type Builder struct {
	RichErr
}

func New(op Op) Builder {
	return Builder{RichErr{op: op}}
}

func (r Builder) WithOp(op Op) Builder {
	r.op = op

	return r
}

func (r Builder) WithErr(err error) Builder {
	r.wrappedError = err

	return r
}

func (r Builder) WithMessage(message string) Builder {
	r.message = message

	return r
}

func (r Builder) WithKind(kind Kind) Builder {
	r.kind = kind

	return r
}

func (r Builder) WithMeta(meta map[string]interface{}) Builder {
	r.meta = meta

	return r
}

func (r Builder) Error() string {
	if r.message == "" && r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return r.message
}

func (r Builder) Op() Op {
	if r.op != "" {
		return r.op
	}

	var re Builder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return ""
	}

	return re.Op()
}

func (r Builder) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	var re Builder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r Builder) WrappedError() error {
	if r.wrappedError != nil {
		return r.wrappedError
	}

	var re Builder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return nil
	}

	return re.WrappedError()
}

func (r Builder) Message() string {
	if r.message != "" {
		return r.message
	}

	var re Builder
	ok := errors.As(r.wrappedError, &re)
	if ok {
		return re.Message()
	}

	if r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return ""
}

func (r Builder) Meta() map[string]interface{} {
	if len(r.meta) != 0 {
		return r.meta
	}

	var re Builder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return make(map[string]interface{})
	}

	return re.Meta()
}

func (r Builder) Build() RichErr {
	return r.RichErr
}

func Analysis(err error) RichErr {
	var richError Builder
	switch {
	case errors.As(err, &richError):
		var re Builder
		errors.As(err, &re)

		return RichErr{
			op:           re.Op(),
			wrappedError: re.WrappedError(),
			message:      re.Message(),
			kind:         re.Kind(),
			meta:         re.Meta(),
		}

	default:
		return RichErr{}
	}
}
