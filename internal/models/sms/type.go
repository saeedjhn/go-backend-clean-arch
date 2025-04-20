package sms

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type TypeName string

const (
	OTPTypeName    TypeName = "otp"
	ReportTypeName TypeName = "report"
)

var typeNameStrings = map[TypeName]string{ //nolint:gochecknoglobals // nothing
	OTPTypeName:    "otp",
	ReportTypeName: "report",
}

func (a TypeName) IsValidTypeName() bool {
	_, ok := typeNameStrings[a]

	return ok
}

type Type struct {
	ID          types.ID
	Name        TypeName // otp", "report", "marketing", etc
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
