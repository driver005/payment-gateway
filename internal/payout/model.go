package payout

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/balance"
	"github.com/google/uuid"
)

// type PayoutDestination struct {
// 	core.Model

// 	BankAccount *bank.BankAccount
// 	Card        *card.Card
// }

type Payout struct {
	core.Model

	Amount      int    `json:"amount,omitempty"`
	ArrivalDate int    `json:"arrival_date,omitempty"`
	Automatic   bool   `json:"automatic,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	// Destination               *PayoutDestination          `json:"destination,omitempty" database:"foreignKey:id"`
	FailureCode         string `json:"failure_code,omitempty"`
	FailureMessage      string `json:"failure_message,omitempty"`
	Method              string `json:"method,omitempty"`
	SourceType          string `json:"source_type,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Status              string `json:"status,omitempty"`
	Type                string `json:"type,omitempty"`

	BalanceTransactionId        uuid.UUID                   `json:"balance_transaction_id,omitempty"`
	BalanceTransaction          *balance.BalanceTransaction `json:"balance_transaction,omitempty" database:"foreignKey:balance_transaction_id"`
	FailureBalanceTransactionId uuid.UUID                   `json:"failure_balance_transaction_id,omitempty"`
	FailureBalanceTransaction   *balance.BalanceTransaction `json:"failure_balance_transaction,omitempty" database:"foreignKey:failure_balance_transaction_id"`
	OriginalPayoutId            *Payout                     `json:"original_payout_id,omitempty"`
	OriginalPayout              *Payout                     `json:"original_payout,omitempty" database:"foreignKey:original_payout_id"`
	ReversedById                *Payout                     `json:"reversed_by_id,omitempty"`
	ReversedBy                  *Payout                     `json:"reversed_by,omitempty" database:"foreignKey:reversed_by_id"`
}
