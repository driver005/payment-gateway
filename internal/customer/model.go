package customer

import (
	"github.com/driver005/gateway/billing/invoice/settings"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/address"
	"github.com/lib/pq"
)

// CustomerTax
type CustomerTax struct {
	core.Model

	// Surfaces if automatic tax computation is possible given the current customer location information.
	AutomaticTax string `json:"automatic_tax,omitempty"`
	// A recent IP address of the customer used for tax reporting and tax location inference.
	IpAddress string `json:"ip_address,omitempty"`
	// The customer's country as identified by Stripe Tax.
	Country string `json:"country,omitempty"`
	// The data source used to infer the customer's location.
	Source string `json:"source,omitempty"`
	// The customer's state, county, province, or region as identified by Stripe Tax.
	State string `json:"state,omitempty"`
}

type Customer struct {
	core.Model

	Address *address.Address `json:"address,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Balance int              `json:"balance,omitempty"`
	// CashBalance          *CashBalance                     `json:"cash_balance,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Delinquent  bool   `json:"delinquent,omitempty"`
	Description string `json:"description,omitempty"`
	// Discount             *Discount                        `json:"discount,omitempty"`
	Email                string                                  `json:"email,omitempty"`
	InvoiceCreditBalance core.JSONB                              `json:"invoice_credit_balance,omitempty"`
	InvoicePrefix        string                                  `json:"invoice_prefix,omitempty"`
	InvoiceSettings      *settings.InvoiceSettingCustomerSetting `json:"invoice_settings,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Name                 string                                  `json:"name,omitempty"`
	NextInvoiceSequence  int                                     `json:"next_invoice_sequence,omitempty"`
	Phone                string                                  `json:"phone,omitempty"`
	PreferredLocales     pq.StringArray                          `json:"preferred_locales,omitempty" database:"type:varchar(64)[]"`
	Shipping             *address.Shipping                       `json:"shipping,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Sources              *ApmsSourcesSourceList                  `json:"sources,omitempty"`
	// Subscriptions        *SubscriptionList                       `json:"subscriptions,omitempty"`
	Tax       *CustomerTax `json:"tax,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	TaxExempt TaxExempt    `json:"tax_exempt,omitempty"`
	// TaxIds               *TaxIDsList                             `json:"tax_ids,omitempty"`
}
