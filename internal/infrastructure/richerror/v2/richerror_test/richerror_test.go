package richerror_test

import (
	"errors"
	"github.com/kr/pretty"
	"github.com/rotisserie/eris"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror/v2"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/prettyprint"
	"testing"
)

func TestRich(t *testing.T) {
	t.Log("Test Rich")

	richerror.New().
		WithOp("OPERATION").
		WithErr("BAD-REQUEST").
		WithWrapperErr(errors.New("external error: database").Error()).
		WithKind(kind.KindStatusBadRequest).
		WithMeta(map[string]interface{}{
			"timestamp": "2024-10-30T12:00:00Z",
		}).
		WithSource(richerror.Pointer, "username")
	//WithTrace()

	//pretty.Log(r)

	rr := richerror.New().
		//WithErr("new-error").
		WithWrapperErr(errors.New("external3 error: database").Error()).
		WithTrace(true)

	a := richerror.Analysis(rr)

	//prettyprint.PrettyPrintData(rr)

	pretty.Log(
		a,
	)

}

func TestEris(t *testing.T) {
	e := errors.New("database connection failed") // external

	ee := eris.Wrap(e, "service failed to process business logic")
	//eee := eris.WithWrapErr(ee, "handler failed to handle request")

	//t.Error(eee)
	prettyprint.PrettyPrintData(eris.ToJSON(ee, true))
}
