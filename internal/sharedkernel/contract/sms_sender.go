package contract

import "time"

// type SMSStatus string
//
// const (
// 	StatusQueued    SMSStatus = "queued"
// 	StatusSent      SMSStatus = "sent"
// 	StatusDelivered SMSStatus = "delivered"
// 	StatusFailed    SMSStatus = "failed"
// 	StatusBlocked   SMSStatus = "blocked"
// 	StatusRejected  SMSStatus = "rejected"
// )

type Status struct {
	RecID      int
	StatusCode int
	StatusText string
	// SentAt      *time.Time // Optional
	// DeliveredAt *time.Time // Optional
	// Error       string     // Optional
}

//go:generate mockery --name SMSSender
type SMSSender interface {
	SendSingle(receptor string, message string) (int, error)
	SendSingleAt(receptor string, message string, duration time.Duration) (int, error)
	SendBulk(receptors []string, message string) ([]int, error)
	GetStatus(recIDs []int) ([]Status, error)
}
