package richerror

import (
	"errors"
	"github.com/rotisserie/eris"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"sync"
)

type SourceType string

const (
	Pointer   SourceType = "pointer"
	Parameter SourceType = "parameter"
	Header    SourceType = "header"
)

// The singleton instance and a sync.Once to ensure it's only created once
var instance *RichError
var once sync.Once

type RichError struct {
	kind    kind.Kind // Status
	source  map[SourceType]string
	meta    map[string]interface{} // meta
	rootErr error
	wrapErr error
}

func New() *RichError {
	once.Do(func() {
		instance = &RichError{
			source: make(map[SourceType]string),
			meta:   make(map[string]interface{}),
		}
	})

	return instance
}

func (r *RichError) WithErr(msg string) *RichError {
	e := eris.New(msg)

	r.rootErr = e
	r.wrapErr = e

	return r
}

func (r *RichError) WithWrapperErr(msg string) *RichError {
	r.wrapErr = eris.Wrap(r.wrapErr, msg)

	return r
}

func (r *RichError) WithKind(kind kind.Kind) *RichError {
	r.kind = kind

	return r
}

func (r *RichError) WithMeta(meta map[string]interface{}) *RichError {
	for k, v := range meta {
		r.meta[k] = v
	}

	return r
}

func (r *RichError) WithSource(t SourceType, v string) *RichError {
	r.source[t] = v

	return r
}

func (r *RichError) WithTrace(pretty bool) *RichError {

	r.meta["trace"] = eris.ToJSON(r.wrapErr, pretty)

	return r
}

func (r *RichError) Error() string {
	return r.rootErr.Error()
}

func (r *RichError) ErrorWithWrap() string {
	return eris.ToCustomString(r.wrapErr, eris.NewDefaultStringFormat(eris.FormatOptions{
		InvertOutput: true, // Invert order to show latest errors first
	}))
}

func (r *RichError) Kind() kind.Kind {
	return r.kind
}

func (r *RichError) Meta() map[string]interface{} {
	meta := r.meta
	r.meta = make(map[string]interface{})

	return meta
}

func (r *RichError) Source() map[SourceType]string {
	return r.source
}

func Analysis(err error) *RichError {
	var re *RichError

	if errors.As(err, &re) {
		return &RichError{
			rootErr: re.rootErr,
			wrapErr: re.wrapErr,
			kind:    re.Kind(),
			source:  re.Source(),
			meta:    re.Meta(),
		}
	}

	return nil
}
