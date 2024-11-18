package richerror

import (
	"errors"
	"maps"

	"github.com/rotisserie/eris"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
)

type SourceType string
type Op string

const (
	Pointer   SourceType = "pointer"
	Parameter SourceType = "parameter"
	Header    SourceType = "header"
)

var metaMap = make(map[string]interface{})

type RichError struct {
	op     Op
	kind   kind.Kind
	source map[SourceType]string
	meta   map[string]interface{}
	rErr   error
	wErr   error
}

func New() *RichError {
	return &RichError{
		source: make(map[SourceType]string),
		meta:   make(map[string]interface{}),
	}
}

func (r *RichError) WithOp(op Op) *RichError {
	r.op = op

	return r
}

func (r *RichError) WithErr(err error) *RichError {
	r.rErr = err
	//r.wErr = eris.New(err.Error())

	return r
}

func (r *RichError) WithWrapErr(err error, msg string) *RichError {
	if err == nil || r.rErr != nil {
		return r
	}

	var re *RichError
	ie := err

	r.rErr = err
	if errors.As(err, &re) {

		if re.wErr != nil {
			ie = re.wErr
		}
	}

	r.wErr = eris.Wrap(ie, msg)

	return r
}

func (r *RichError) WithKind(kind kind.Kind) *RichError {
	r.kind = kind

	return r
}

func (r *RichError) WithMeta(meta map[string]interface{}) *RichError {
	//for k, v := range meta {
	//	r.meta[k] = v
	//}
	maps.Copy(r.meta, meta)

	return r
}

func (r *RichError) WithSource(t SourceType, v string) *RichError {
	r.source[t] = v

	return r
}

func (r *RichError) WithTrace(pretty bool) *RichError {
	r.meta["trace"] = eris.ToJSON(r.wErr, pretty)

	return r
}

func (r *RichError) Op() Op {
	if len(r.op) != 0 {
		return r.op
	}

	var re *RichError

	if errors.As(r.rErr, &re) {
		return re.Op()
	}

	return ""
}

func (r *RichError) Kind() kind.Kind {
	if r.kind != 0 {
		return r.kind
	}

	var re *RichError

	if errors.As(r.rErr, &re) {
		return re.Kind()
	}

	return 0
}

func (r *RichError) Meta() map[string]interface{} {
	var re *RichError

	maps.Copy(metaMap, r.meta)

	if errors.As(r.rErr, &re) {
		maps.Copy(metaMap, re.meta)

		return re.Meta()
	}

	return metaMap
}

func (r *RichError) Source() map[SourceType]string {
	if len(r.source) != 0 {
		return r.source
	}

	var re *RichError

	if errors.As(r.rErr, &re) {
		return re.Source()
	}

	return nil
}

func (r *RichError) Error() string {
	if r.rErr != nil {
		return r.rErr.Error()
	}

	return "nil"
}

func (r *RichError) ErrorWithWrap() string {
	return eris.ToCustomString(r.wErr, eris.NewDefaultStringFormat(eris.FormatOptions{
		InvertOutput: true, // Invert order to show latest errors first
	}))
}

func Analysis(err error) *RichError {
	var re *RichError

	if errors.As(err, &re) {
		return &RichError{
			op:     re.Op(),
			kind:   re.Kind(),
			source: re.Source(),
			meta:   re.Meta(),
			rErr:   re.rErr,
			wErr:   re.wErr,
		}
	}

	return nil
}
