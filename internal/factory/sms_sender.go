package factory

import (
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/kavenegar"
	"github.com/saeedjhn/go-backend-clean-arch/internal/models/sms"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

var ErrUnsupportedProvider = errors.New("unsupported SMS provider")

func SMSSenderFactory(provider sms.ProviderName) (contract.SMSSender, error) {
	if !provider.IsValidProviderName() {
		return nil, ErrUnsupportedProvider
	}

	switch provider { //nolint:exhaustive // nothing
	case sms.Kavenegar:
		return kavenegar.New(), nil
	// case sms.MelliPayamak:
	// case sms.SMSIR:
	// case sms.FarazSMS:
	// case sms.Payamito:
	// case sms.Farapayamak:
	default:
		return kavenegar.New(), nil
	}
}
