package customer

import (
	"github.com/driver005/gateway/billing/invoice/settings"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/utils/address"
)

// CustomerDefaultSource ID of the default payment source for the customer.  If you are using payment methods created via the PaymentMethods API, see the [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/object#customer_object-invoice_settings-default_payment_method) field instead.
type CustomerDefaultSource struct {
	// BankAccount *BankAccount
	// Card        *Card
	// Source      *Source
}

// CustomerTaxLocation
type CustomerTaxLocation struct {
	// The customer's country as identified by Stripe Tax.
	Country string `json:"country"`
	// The data source used to infer the customer's location.
	Source string `json:"source"`
	// The customer's state, county, province, or region as identified by Stripe Tax.
	State string `json:"state,omitempty"`
}

// CustomerTax
type CustomerTax struct {
	// Surfaces if automatic tax computation is possible given the current customer location information.
	AutomaticTax string `json:"automatic_tax"`
	// A recent IP address of the customer used for tax reporting and tax location inference.
	IpAddress string              `json:"ip_address,omitempty"`
	Location  CustomerTaxLocation `json:"location,omitempty"`
}

type Customer struct {
	Address address.Address `json:"address,omitempty"`
	Balance int             `json:"balance,omitempty"`
	// CashBalance          *CashBalance                     `json:"cash_balance,omitempty"`
	Created       int                   `json:"created"`
	Currency      models.Currency       `json:"currency,omitempty"`
	DefaultSource CustomerDefaultSource `json:"default_source,omitempty"`
	Delinquent    bool                  `json:"delinquent,omitempty"`
	Description   string                `json:"description,omitempty"`
	// Discount             *Discount                        `json:"discount,omitempty"`
	Email                string                                  `json:"email,omitempty"`
	Id                   string                                  `json:"id"`
	InvoiceCreditBalance *map[string]int                         `json:"invoice_credit_balance,omitempty"`
	InvoicePrefix        string                                  `json:"invoice_prefix,omitempty"`
	InvoiceSettings      *settings.InvoiceSettingCustomerSetting `json:"invoice_settings,omitempty"`
	Livemode             bool                                    `json:"livemode"`
	Metadata             *map[string]string                      `json:"metadata,omitempty"`
	Name                 string                                  `json:"name,omitempty"`
	NextInvoiceSequence  int                                     `json:"next_invoice_sequence,omitempty"`
	Object               string                                  `json:"object"`
	Phone                string                                  `json:"phone,omitempty"`
	PreferredLocales     []string                                `json:"preferred_locales,omitempty"`
	Shipping             address.Shipping                        `json:"shipping,omitempty"`
	// Sources              *ApmsSourcesSourceList                  `json:"sources,omitempty"`
	// Subscriptions        *SubscriptionList                       `json:"subscriptions,omitempty"`
	Tax       *CustomerTax `json:"tax,omitempty"`
	TaxExempt string       `json:"tax_exempt,omitempty"`
	// TaxIds               *TaxIDsList                             `json:"tax_ids,omitempty"`
}
