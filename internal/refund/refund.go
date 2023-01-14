package refund

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/models"
)

// RefundNextAction
type RefundNextAction struct {
	EmailSent models.EmailSent `json:"email_sent"`
	// The expiry timestamp.
	ExpiresAt int `json:"expires_at"`
	// Type of the next action to perform.
	Type string `json:"type"`
}

type Refund struct {
	Amount             int                         `json:"amount"`
	BalanceTransaction *balance.BalanceTransaction `json:"balance_transaction,omitempty"`
	// Charge             *charge.Charge              `json:"charge,omitempty"`
	Created                   int                         `json:"created"`
	Currency                  string                      `json:"currency"`
	Description               *string                     `json:"description,omitempty"`
	FailureBalanceTransaction *balance.BalanceTransaction `json:"failure_balance_transaction,omitempty"`
	FailureReason             *string                     `json:"failure_reason,omitempty"`
	Id                        string                      `json:"id"`
	InstructionsEmail         *string                     `json:"instructions_email,omitempty"`
	Metadata                  core.JSONB                  `json:"metadata,omitempty"`
	NextAction                *RefundNextAction           `json:"next_action,omitempty"`
	Object                    string                      `json:"object"`
	// PaymentIntent intent.PaymentIntent `json:"payment_intent,omitempty"`
	Reason        string `json:"reason,omitempty"`
	ReceiptNumber string `json:"receipt_number,omitempty"`
	Status        string `json:"status,omitempty"`
}
