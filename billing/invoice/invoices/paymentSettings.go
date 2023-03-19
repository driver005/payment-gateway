package invoices

import (
	"github.com/driver005/gateway/billing/invoice/options"
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// InvoicesPaymentSettings
type InvoicesPaymentSettings struct {
	core.Model
	// ID of the mandate to be used for this invoice. It must correspond to the payment method used to pay the invoice, including the invoice's default_payment_method or default_source, if set.
	DefaultMandate       string                               `json:"default_mandate,omitempty"`
	PaymentMethodOptions *options.InvoicePaymentMethodOptions `json:"payment_method_options,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// The list of payment method types (e.g. card) to provide to the invoice’s PaymentIntent. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice’s default payment method, the subscription’s default payment method, the customer’s default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
	PaymentMethodTypes pq.StringArray `json:"payment_method_types,omitempty" database:"type:varchar(64)[]"`
}
