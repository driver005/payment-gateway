package card

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/utils/account"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Card struct {
	core.Model

	AddressCity            string             `json:"address_city,omitempty"`
	AddressCountry         string             `json:"address_country,omitempty"`
	AddressLine1           string             `json:"address_line1,omitempty"`
	AddressLine1Check      AddressLine1Check  `json:"address_line1_check,omitempty"`
	AddressLine2           string             `json:"address_line2,omitempty"`
	AddressState           string             `json:"address_state,omitempty"`
	AddressZip             string             `json:"address_zip,omitempty"`
	AddressZipCheck        AddressZipCheck    `json:"address_zip_check,omitempty"`
	AvailablePayoutMethods pq.StringArray     `json:"available_payout_methods,omitempty" database:"type:varchar(64)[]"`
	Brand                  Brand              `json:"brand"`
	Country                string             `json:"country,omitempty"`
	Currency               string             `json:"currency,omitempty"`
	CvcCheck               CvcCheck           `json:"cvc_check,omitempty"`
	DefaultForCurrency     bool               `json:"default_for_currency,omitempty"`
	DynamicLast4           string             `json:"dynamic_last4,omitempty"`
	ExpMonth               int                `json:"exp_month"`
	ExpYear                int                `json:"exp_year"`
	Fingerprint            string             `json:"fingerprint,omitempty"`
	Funding                Funding            `json:"funding"`
	Last4                  string             `json:"last4"`
	Name                   string             `json:"name,omitempty"`
	Status                 string             `json:"status,omitempty"`
	TokenizationMethod     TokenizationMethod `json:"tokenization_method,omitempty"`

	AccountId  *uuid.UUID         `json:"account_id,omitempty" swaggerignore:"true"`
	Account    *account.Account   `json:"account,omitempty" database:"foreignKey:account_id"`
	CustomerId *uuid.UUID         `json:"customer_id" swaggerignore:"true"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
