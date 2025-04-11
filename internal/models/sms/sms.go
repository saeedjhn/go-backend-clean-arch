package sms

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type Config struct {
	ID           types.ID
	Title        string
	IsDefault    bool
	Priority     int
	Credentials  Credential
	Status       Status
	ProviderID   types.ID
	SenderLineID types.ID
	Type         types.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Provider struct {
	ID          types.ID
	Name        string
	Slug        string
	Description string
	Website     string
	APIPaths    APIPath
	Status      Status
	SenderLines []types.ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SenderLine struct {
	ID          types.ID
	Number      string
	Capacity    int
	IsActive    bool
	Description string
	ProviderID  types.ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Type struct {
	ID          types.ID
	Name        string // otp", "report", "marketing", etc
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Credential struct {
	APIKey    string
	SecretKey string
	Username  string
	Password  string
}

type APIPath struct {
	BaseURL             string
	SendAPIPath         *string
	StatusAPIPath       *string
	BalanceAPIPath      *string
	ReportAPIPath       *string
	SenderLinesAPIPath  *string
	TemplateSendAPIPath *string
	BulkSendAPIPath     *string
	InboxAPIPath        *string
}
