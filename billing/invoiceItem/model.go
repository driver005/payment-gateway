package invoiceItem

import (
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/models"
)

// Invoiceitem Sometimes you want to add a charge or credit to a customer, but actually charge or credit the customer's card only at the end of a regular billing cycle. This is useful for combining several charges (to minimize per-transaction fees), or for having Stripe tabulate your usage-based billing totals.  Related guide: [Subscription Invoices](https://stripe.com/docs/billing/invoices/subscription#adding-upcoming-invoice-items).
type Invoiceitem struct {
	// Amount (in the `currency` specified) of the invoice item. This should always be equal to `unit_amount * quantity`.
	Amount int `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string              `json:"currency"`
	Customer customer.Customer `json:"customer"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Date int `json:"date"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// If true, discounts will apply to this invoice item. Always false for prorations.
	Discountable bool `json:"discountable"`
	// The discounts which apply to the invoice item. Item discounts are applied before invoice discounts. Use `expand[]=discounts` to expand each discount.
	// Discounts []Discount `json:"discounts,omitempty"`
	// Unique identifier for the object.
	Id      string             `json:"id"`
	Invoice invoice.Invoice `json:"invoice,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string                        `json:"object"`
	Period invoice.InvoiceLineItemPeriod `json:"period"`
	// Price  *Price              `json:"price,omitempty"`
	// Whether the invoice item was created automatically as a proration adjustment when the customer switched plans.
	Proration bool `json:"proration"`
	// Quantity of units for the invoice item. If the invoice item is a proration, the quantity of the subscription that the proration was computed for.
	Quantity     int                     `json:"quantity"`
	// Subscription Subscription `json:"subscription,omitempty"`
	// The subscription item that this invoice item has been created for, if any.
	SubscriptionItem string `json:"subscription_item,omitempty"`
	// The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item.
	TaxRates []models.TaxRate `json:"tax_rates,omitempty"`
	// Unit amount (in the `currency` specified) of the invoice item.
	UnitAmount int `json:"unit_amount,omitempty"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,omitempty"`
}
