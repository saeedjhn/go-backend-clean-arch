package httpstatus //nolint:cyclop // the average complexity for the package httpstatus is 24.000000, max is 10.000000

import (
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
)

func MapkindToHTTPStatusCode(k kind.Kind) int {
	switch k { //nolint:exhaustive // missing cases in switch of type kind.Kind
	case kind.KindStatusContinue:
		return http.StatusContinue
	case kind.KindStatusSwitchingProtocols:
		return http.StatusSwitchingProtocols
	case kind.KindStatusProcessing:
		return http.StatusProcessing
	case kind.KindStatusEarlyHints:
		return http.StatusEarlyHints

	case kind.KindStatusOK:
		return http.StatusOK
	case kind.KindStatusCreated:
		return http.StatusCreated
	case kind.KindStatusAccepted:
		return http.StatusAccepted
	case kind.KindStatusNonAuthoritativeInfo:
		return http.StatusNonAuthoritativeInfo
	case kind.KindStatusNoContent:
		return http.StatusNoContent
	case kind.KindStatusResetContent:
		return http.StatusResetContent
	case kind.KindStatusPartialContent:
		return http.StatusPartialContent
	case kind.KindStatusMultiStatus:
		return http.StatusMultiStatus
	case kind.KindStatusAlreadyReported:
		return http.StatusAlreadyReported
	case kind.KindStatusIMUsed:
		return http.StatusIMUsed

	case kind.KindStatusBadRequest:
		return http.StatusBadRequest
	case kind.KindStatusUnauthorized:
		return http.StatusUnauthorized
	case kind.KindStatusPaymentRequired:
		return http.StatusPaymentRequired
	case kind.KindStatusForbidden:
		return http.StatusForbidden
	case kind.KindStatusNotFound:
		return http.StatusNotFound
	case kind.KindStatusConflict:
		return http.StatusConflict
	case kind.KindStatusUnprocessableEntity:
		return http.StatusUnprocessableEntity

	case kind.KindStatusInternalServerError:
		return http.StatusInternalServerError

	default:
		return http.StatusBadRequest
	}
}
