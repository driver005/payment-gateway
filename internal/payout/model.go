package payout

import (
	"github.com/driver005/gateway/internal/balance"
)

type PayoutDestination struct {
	// BankAccount        *BankAccount
	// Card               *Card
	// DeletedBankAccount *DeletedBankAccount
	// DeletedCard        *DeletedCard
}

type Payout struct {
	Amount int `json:"amount"`
	ArrivalDate int `json:"arrival_date"`
	Automatic          bool                       `json:"automatic"`
	BalanceTransaction *balance.BalanceTransaction `json:"balance_transaction,omitempty"`
	Created int `json:"created"`
	Currency string `json:"currency"`
	Description               string                     `json:"description,omitempty"`
	Destination               *PayoutDestination         `json:"destination,omitempty"`
	FailureBalanceTransaction balance.BalanceTransaction `json:"failure_balance_transaction,omitempty"`
	FailureCode string `json:"failure_code,omitempty"`
	FailureMessage string `json:"failure_message,omitempty"`
	Id string `json:"id"`
	Livemode bool `json:"livemode"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Method string `json:"method"`
	Object         string  `json:"object"`
	OriginalPayout *Payout `json:"original_payout,omitempty"`
	ReversedBy     *Payout `json:"reversed_by,omitempty"`
	SourceType string `json:"source_type"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Status string `json:"status"`
	Type string `json:"type"`
}
