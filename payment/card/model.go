package card

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/utils/account"
	"github.com/google/uuid"
)

type Card struct {
	core.Model

	AddressCity            string   `json:"address_city,omitempty"`
	AddressCountry         string   `json:"address_country,omitempty"`
	AddressLine1           string   `json:"address_line1,omitempty"`
	AddressLine1Check      string   `json:"address_line1_check,omitempty"`
	AddressLine2           string   `json:"address_line2,omitempty"`
	AddressState           string   `json:"address_state,omitempty"`
	AddressZip             string   `json:"address_zip,omitempty"`
	AddressZipCheck        string   `json:"address_zip_check,omitempty"`
	AvailablePayoutMethods []string `json:"available_payout_methods,omitempty" database:"type:text[]"`
	Brand                  string   `json:"brand"`
	Country                string   `json:"country,omitempty"`
	Currency               string   `json:"currency,omitempty"`
	CvcCheck               string   `json:"cvc_check,omitempty"`
	DefaultForCurrency     bool     `json:"default_for_currency,omitempty"`
	DynamicLast4           string   `json:"dynamic_last4,omitempty"`
	ExpMonth               int      `json:"exp_month"`
	ExpYear                int      `json:"exp_year"`
	Fingerprint            string   `json:"fingerprint,omitempty"`
	Funding                string   `json:"funding"`
	Last4                  string   `json:"last4"`
	Name                   string   `json:"name,omitempty"`
	Status                 string   `json:"status,omitempty"`
	TokenizationMethod     string   `json:"tokenization_method,omitempty"`

	AccountId  uuid.UUID          `json:"account_id,omitempty"`
	Account    *account.Account   `json:"account,omitempty" database:"foreignKey:account_id"`
	CustomerId uuid.UUID          `json:"customer_id"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
