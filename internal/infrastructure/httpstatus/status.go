package httpstatus

import (
	. "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
	"net/http"
)

func FromKind(kind Kind) int {
	switch kind {
	case KindStatusContinue:
		return http.StatusContinue
	case KindStatusSwitchingProtocols:
		return http.StatusSwitchingProtocols
	case KindStatusProcessing:
		return http.StatusProcessing
	case KindStatusEarlyHints:
		return http.StatusEarlyHints

	case KindStatusOK:
		return http.StatusOK
	case KindStatusCreated:
		return http.StatusCreated
	case KindStatusAccepted:
		return http.StatusAccepted
	case KindStatusNonAuthoritativeInfo:
		return http.StatusNonAuthoritativeInfo
	case KindStatusNoContent:
		return http.StatusNoContent
	case KindStatusResetContent:
		return http.StatusResetContent
	case KindStatusPartialContent:
		return http.StatusPartialContent
	case KindStatusMultiStatus:
		return http.StatusMultiStatus
	case KindStatusAlreadyReported:
		return http.StatusAlreadyReported
	case KindStatusIMUsed:
		return http.StatusIMUsed

	case KindStatusBadRequest:
		return http.StatusBadRequest
	case KindStatusUnauthorized:
		return http.StatusUnauthorized
	case KindStatusPaymentRequired:
		return http.StatusPaymentRequired
	case KindStatusForbidden:
		return http.StatusForbidden
	case KindStatusNotFound:
		return http.StatusNotFound
	case KindStatusConflict:
		return http.StatusConflict
	case KindStatusUnprocessableEntity:
		return http.StatusUnprocessableEntity

	case KindStatusInternalServerError:
		return http.StatusInternalServerError

	default:
		return http.StatusBadRequest
	}
}
