package models

import (
	"github.com/driver005/gateway/core"
	"github.com/gofrs/uuid"
)

// Refund - Refund represent an amount of money transfered back to the Customer for a given reason. Refunds may occur in relation to Returns, Swaps and Claims, but can also be initiated by a store operator.
type Refund struct {
	core.Model

	// The id of the Order that the Refund is related to.
	OrderId uuid.NullUUID `json:"order_id"`

	// The amount that has be refunded to the Customer.
	Amount int32 `json:"amount"`

	// An optional note explaining why the amount was refunded.
	Note string `json:"note" database:"default:null"`

	// The reason given for the Refund, will automatically be set when processed as part of a Swap, Claim or Return.
	Reason string `json:"reason" database:"default:null"`

	// Randomly generated key used to continue the completion of the refund in case of failure.
	IdempotencyKey string `json:"idempotency_key" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
