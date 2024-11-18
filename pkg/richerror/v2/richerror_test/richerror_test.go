package richerror_test

import (
	"errors"
	"testing"

	"github.com/rotisserie/eris"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/prettyprint"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror/v2"
)

func TestRich(t *testing.T) {
	t.Log("Test Rich")
	e := errors.New("database: row not found")

	r1 := richerror.New().
		WithOp("OPERATION").
		//WithErr(e).
		WithWrapErr(e, "repository - record not found").
		WithKind(kind.KindStatusBadRequest).
		WithMeta(map[string]interface{}{
			"timestamp": "2024-10-30T12:00:00Z",
		}).
		WithSource(richerror.Pointer, "username")

	r2 := richerror.New().
		WithWrapErr(r1, "usecase - error wrapper with repository").
		WithMeta(map[string]interface{}{
			"query": "QUERY",
		})

	r3 := richerror.New().
		//WithErr(errors.New("AAAA")).
		WithWrapErr(r2, "handler - error wrapper with usecase").
		WithMeta(map[string]interface{}{
			"db": "DB",
		}).
		WithTrace(false)

	prettyprint.Log(r3.Op())
	prettyprint.Log(r3.Kind())
	prettyprint.Log(r3.Source())
	prettyprint.Log(r3.ErrorWithWrap())
	prettyprint.Log(r3.Error())
	prettyprint.Log(r3.Meta())

	//
	//a := richerror.Analysis(r2)
	//
	//prettyprint.Log(a)
	//
	//pretty.Log(
	//	a,
	//)

}

func TestEris(t *testing.T) {
	e := errors.New("database connection failed") // external

	ee := eris.Wrap(e, "service failed to process business logic")
	//eee := eris.WithWrapErr(ee, "handler failed to handle request")

	//t.Error(eee)
	prettyprint.Log(eris.ToJSON(ee, true))
}
