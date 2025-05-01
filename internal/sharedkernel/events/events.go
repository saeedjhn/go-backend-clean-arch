package events

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
)

const (
	UsersRegistered = contract.Topic("users.registered")
	// PurchaseSucceedTopic = contract.Topic("payment.purchase_succeed").
	// PurchaseFailedTopic  = contract.Topic("payment.purchase_failed").
)
