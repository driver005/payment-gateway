package bank

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/utils/account"
	"github.com/google/uuid"
)

type BankAccount struct {
	core.Model

	AccountHolderName      string   `json:"account_holder_name,omitempty"`
	AccountHolderType      string   `json:"account_holder_type,omitempty"`
	AccountType            string   `json:"account_type,omitempty"`
	AvailablePayoutMethods []string `json:"available_payout_methods,omitempty" database:"type:text[]"`
	BankName               string   `json:"bank_name,omitempty"`
	Country                string   `json:"country"`
	Currency               string   `json:"currency"`
	DefaultForCurrency     bool     `json:"default_for_currency,omitempty"`
	Fingerprint            string   `json:"fingerprint,omitempty"`
	Last4                  string   `json:"last4"`
	RoutingNumber          string   `json:"routing_number,omitempty"`
	Status                 string   `json:"status"`

	AccountId  uuid.NullUUID      `json:"account_id,omitempty"`
	Account    *account.Account   `json:"account,omitempty" database:"foreignKey:account_id"`
	CustomerId uuid.NullUUID      `json:"customer_id,omitempty"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
