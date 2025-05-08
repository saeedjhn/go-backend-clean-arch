package events

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

const (
	UsersRegistered = models.EventType("users.registered")
	// PurchaseSucceedTopic = contract.EventType("payment.purchase_succeed").
	// PurchaseFailedTopic  = contract.EventType("payment.purchase_failed").
)
