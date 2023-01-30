package refund

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/utils/email"
	"github.com/google/uuid"
)

// RefundNextAction
type RefundNextAction struct {
	core.Model

	EmailSent email.EmailSent `json:"email_sent,omitempty" database:"foreignKey:id"`
	// The expiry timestamp.
	ExpiresAt int `json:"expires_at,omitempty"`
	// Type of the next action to perform.
	Type string `json:"type,omitempty"`
}

type Refund struct {
	core.Model

	Amount            int               `json:"amount,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	Description       string            `json:"description,omitempty"`
	FailureReason     string            `json:"failure_reason,omitempty"`
	InstructionsEmail string            `json:"instructions_email,omitempty"`
	NextAction        *RefundNextAction `json:"next_action,omitempty" database:"foreignKey:id"`
	Reason            string            `json:"reason,omitempty"`
	ReceiptNumber     string            `json:"receipt_number,omitempty"`
	Status            string            `json:"status,omitempty"`

	ChargeId                    uuid.UUID                   `json:"charge_id,omitempty"`
	Charge                      *charge.Charge              `json:"charge,omitempty" database:"foreignKey:charge_id"`
	PaymentIntentId             uuid.UUID                   `json:"payment_intent_id,omitempty"`
	PaymentIntent               *intent.PaymentIntent       `json:"payment_intent,omitempty" database:"foreignKey:payment_intent_id"`
	BalanceTransactionId        uuid.UUID                   `json:"balance_transaction_id,omitempty"`
	BalanceTransaction          *balance.BalanceTransaction `json:"balance_transaction,omitempty" database:"foreignKey:balance_transaction_id"`
	FailureBalanceTransactionId uuid.UUID                   `json:"failure_balance_transaction_id,omitempty"`
	FailureBalanceTransaction   *balance.BalanceTransaction `json:"failure_balance_transaction,omitempty" database:"foreignKey:failure_balance_transaction_id"`
}
