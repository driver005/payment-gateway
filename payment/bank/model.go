package bank

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/utils/account"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type BankAccount struct {
	core.Model

	AccountHolderName      string            `json:"account_holder_name,omitempty"`
	AccountHolderType      AccountHolderType `json:"account_holder_type,omitempty"`
	AccountType            AccountType       `json:"account_type,omitempty"`
	AvailablePayoutMethods pq.StringArray    `json:"available_payout_methods,omitempty" database:"type:varchar(64)[]"`
	BankName               string            `json:"bank_name,omitempty"`
	Country                string            `json:"country"`
	Currency               string            `json:"currency"`
	DefaultForCurrency     bool              `json:"default_for_currency,omitempty"`
	Fingerprint            string            `json:"fingerprint,omitempty"`
	Last4                  string            `json:"last4"`
	RoutingNumber          string            `json:"routing_number,omitempty"`
	Status                 Status            `json:"status"`

	AccountId  uuid.NullUUID      `json:"account_id,omitempty"`
	Account    *account.Account   `json:"account,omitempty" database:"foreignKey:account_id"`
	CustomerId uuid.NullUUID      `json:"customer_id,omitempty"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
