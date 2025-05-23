package sms

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type ProviderName string

const (
	MelliPayamak ProviderName = "mellipayamak"
	Kavenegar    ProviderName = "kavenegar"
	SMSIR        ProviderName = "smsir"
	FarazSMS     ProviderName = "farazsms"
	Payamito     ProviderName = "payamito"
	Farapayamak  ProviderName = "farapayamak"
	Unknown      ProviderName = "unknown"
)

var providerNameStrings = map[ProviderName]string{ //nolint:gochecknoglobals // nothing
	MelliPayamak: "mellipayamak",
	Kavenegar:    "kavenegar",
	SMSIR:        "smsir",
	FarazSMS:     "farazsms",
	Payamito:     "payamito",
	Farapayamak:  "farapayamak",
	Unknown:      "unknown",
}

func (a ProviderName) IsValidProviderName() bool {
	_, ok := providerNameStrings[a]

	return ok
}

type Provider struct {
	ID          types.ID
	Name        ProviderName
	Slug        string
	Description string
	Website     string
	APIPaths    APIPath
	Credentials Credential
	Status      Status
	SenderLines []types.ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Credential struct {
	APIKey   string
	Username string
	Password string
	// SecretKey string
}

type APIPath struct {
	BaseURL       string
	Single        string
	Bulk          string
	ReceiveStatus string
	// etc...
}

// Factory pattern
// var ErrUnsupportedProvider = errors.New("unsupported SMS provider")
//
// func SMSSenderFactory(provider sms.ProviderName) (contract.SMSSender, error) {
// 	if !provider.IsValidProviderName() {
// 		return nil, ErrUnsupportedProvider
// 	}
//
// 	switch provider { //nolint:exhaustive // nothing
// 	case sms.Kavenegar:
// 		return kavenegar.New(), nil
// 	// case sms.MelliPayamak:
// 	// case sms.SMSIR:
// 	// case sms.FarazSMS:
// 	// case sms.Payamito:
// 	// case sms.Farapayamak:
// 	default:
// 		return kavenegar.New(), nil
// 	}
// }
