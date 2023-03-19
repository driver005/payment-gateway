package payout

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/balance"
	"github.com/google/uuid"
)

type Payout struct {
	core.Model

	Amount              int        `json:"amount,omitempty"`
	ArrivalDate         int        `json:"arrival_date,omitempty"`
	Automatic           bool       `json:"automatic,omitempty"`
	Currency            string     `json:"currency,omitempty"`
	Description         string     `json:"description,omitempty"`
	Destination         string     `json:"destination,omitempty"`
	FailureCode         string     `json:"failure_code,omitempty"`
	FailureMessage      string     `json:"failure_message,omitempty"`
	Method              Method     `json:"method,omitempty"`
	SourceType          SourceType `json:"source_type,omitempty"`
	StatementDescriptor string     `json:"statement_descriptor,omitempty"`
	Status              Status     `json:"status,omitempty"`
	Type                Type       `json:"type,omitempty"`

	BalanceTransactionId        *uuid.UUID                  `json:"balance_transaction_id,omitempty" swaggerignore:"true"`
	BalanceTransaction          *balance.BalanceTransaction `json:"balance_transaction,omitempty" database:"foreignKey:balance_transaction_id" swaggertype:"primitive,string" format:"uuid"`
	FailureBalanceTransactionId *uuid.UUID                  `json:"failure_balance_transaction_id,omitempty" swaggerignore:"true"`
	FailureBalanceTransaction   *balance.BalanceTransaction `json:"failure_balance_transaction,omitempty" database:"foreignKey:failure_balance_transaction_id" swaggertype:"primitive,string" format:"uuid"`
	OriginalPayoutId            *uuid.UUID                  `json:"original_payout_id,omitempty" swaggerignore:"true"`
	OriginalPayout              *Payout                     `json:"original_payout,omitempty" database:"foreignKey:original_payout_id" swaggertype:"primitive,string" format:"uuid"`
	ReversedById                *uuid.UUID                  `json:"reversed_by_id,omitempty" swaggerignore:"true"`
	ReversedBy                  *Payout                     `json:"reversed_by,omitempty" database:"foreignKey:reversed_by_id" swaggertype:"primitive,string" format:"uuid"`
}
