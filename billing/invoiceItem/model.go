package invoiceItem

import (
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/billing/subscriptionItem"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/utils/tax"
	"github.com/google/uuid"
)

// Invoiceitem Sometimes you want to add a charge or credit to a customer, but actually charge or credit the customer's card only at the end of a regular billing cycle. This is useful for combining several charges (to minimize per-transaction fees), or for having Stripe tabulate your usage-based billing totals.  Related guide: [Subscription Invoices](https://stripe.com/docs/billing/invoices/subscription#adding-upcoming-invoice-items).
type Invoiceitem struct {
	core.Model

	// Amount (in the `currency` specified) of the invoice item. This should always be equal to `unit_amount * quantity`.
	Amount int `json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Date int `json:"date,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// If true, discounts will apply to this invoice item. Always false for prorations.
	Discountable bool `json:"discountable,omitempty"`
	// Unique identifier for the object.
	PeriodEnd   int `json:"period_end,omitempty"`
	PeriodStart int `json:"period_start,omitempty"`
	// Whether the invoice item was created automatically as a proration adjustment when the customer switched plans.
	Proration bool `json:"proration,omitempty"`
	// Quantity of units for the invoice item. If the invoice item is a proration, the quantity of the subscription that the proration was computed for.
	Quantity int `json:"quantity,omitempty"`
	// The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item.
	TaxRates *tax.TaxRate `json:"tax_rates,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Unit amount (in the `currency` specified) of the invoice item.
	UnitAmount int `json:"unit_amount,omitempty"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,omitempty"`

	// The discounts which apply to the invoice item. Item discounts are applied before invoice discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []discount.Discount `json:"discounts,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`

	CustomerId         *uuid.UUID                         `json:"customer_id,omitempty" swaggerignore:"true"`
	Customer           *customer.Customer                 `json:"customer,omitempty" database:"foreignKey:customer_id" swaggertype:"primitive,string" format:"uuid"`
	InvoiceId          *uuid.UUID                         `json:"invoice_id,omitempty" swaggerignore:"true"`
	Invoice            *invoice.Invoice                   `json:"invoice,omitempty" database:"foreignKey:invoice_id" swaggertype:"primitive,string" format:"uuid"`
	PriceId            *uuid.UUID                         `json:"price_id,omitempty" swaggerignore:"true"`
	Price              *price.Price                       `json:"price,omitempty" database:"foreignKey:price_id" swaggertype:"primitive,string" format:"uuid"`
	SubscriptionId     *uuid.UUID                         `json:"subscription_id,omitempty" swaggerignore:"true"`
	Subscription       *subscription.Subscription         `json:"subscription,omitempty" database:"foreignKey:subscription_id" swaggertype:"primitive,string" format:"uuid"`
	SubscriptionItemId *uuid.UUID                         `json:"subscription_item_id,omitempty" swaggerignore:"true"`
	SubscriptionItem   *subscriptionItem.SubscriptionItem `json:"subscription_item,omitempty" database:"foreignKey:subscription_item_id" swaggertype:"primitive,string" format:"uuid"`
}
