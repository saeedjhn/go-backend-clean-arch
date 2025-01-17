package httpstatus //nolint:cyclop // the average complexity for the package httpstatus is 24.000000, max is 10.000000

import (
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func MapkindToHTTPStatusCode(k richerror.Kind) int {
	switch k { //nolint:exhaustive // missing cases in switch of type kind.Kind
	case richerror.KindStatusContinue:
		return http.StatusContinue
	case richerror.KindStatusSwitchingProtocols:
		return http.StatusSwitchingProtocols
	case richerror.KindStatusProcessing:
		return http.StatusProcessing
	case richerror.KindStatusEarlyHints:
		return http.StatusEarlyHints

	case richerror.KindStatusOK:
		return http.StatusOK
	case richerror.KindStatusCreated:
		return http.StatusCreated
	case richerror.KindStatusAccepted:
		return http.StatusAccepted
	case richerror.KindStatusNonAuthoritativeInfo:
		return http.StatusNonAuthoritativeInfo
	case richerror.KindStatusNoContent:
		return http.StatusNoContent
	case richerror.KindStatusResetContent:
		return http.StatusResetContent
	case richerror.KindStatusPartialContent:
		return http.StatusPartialContent
	case richerror.KindStatusMultiStatus:
		return http.StatusMultiStatus
	case richerror.KindStatusAlreadyReported:
		return http.StatusAlreadyReported
	case richerror.KindStatusIMUsed:
		return http.StatusIMUsed

	case richerror.KindStatusBadRequest:
		return http.StatusBadRequest
	case richerror.KindStatusUnauthorized:
		return http.StatusUnauthorized
	case richerror.KindStatusPaymentRequired:
		return http.StatusPaymentRequired
	case richerror.KindStatusForbidden:
		return http.StatusForbidden
	case richerror.KindStatusNotFound:
		return http.StatusNotFound
	case richerror.KindStatusConflict:
		return http.StatusConflict
	case richerror.KindStatusUnprocessableEntity:
		return http.StatusUnprocessableEntity

	case richerror.KindStatusInternalServerError:
		return http.StatusInternalServerError

	default:
		return http.StatusBadRequest
	}
}
