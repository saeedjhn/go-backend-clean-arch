package events

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

const (
	UsersRegistered = types.Event("users.registered")
	// PurchaseSucceedTopic = contract.Event("payment.purchase_succeed").
	// PurchaseFailedTopic  = contract.Event("payment.purchase_failed").
)
