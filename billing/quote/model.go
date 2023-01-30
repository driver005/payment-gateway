package quote

import (
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/invoice/settings"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/billing/subscriptionSchedule"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/google/uuid"
)

// QuotesResourceAutomaticTax
type QuotesResourceAutomaticTax struct {
	core.Model
	// Automatically calculate taxes
	Enabled bool `json:"enabled,omitempty"`
	// The status of the most recent automated tax calculation for this quote.
	Status string `json:"status,omitempty"`
}

// QuotesResourceSubscriptionDataSubscriptionData
type QuotesResourceSubscriptionDataSubscriptionData struct {
	core.Model
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description,omitempty"`
	// When creating a new subscription, the date of which the subscription schedule will start after the quote is accepted. This date is ignored if it is in the past when the quote is accepted. Measured in seconds since the Unix epoch.
	EffectiveDate int `json:"effective_date,omitempty"`
	// Integer representing the number of trial period days before the customer is charged for the first time.
	TrialPeriodDays int `json:"trial_period_days,omitempty"`
}

// QuotesResourceStatusTransitions
type QuotesResourceStatusTransitions struct {
	core.Model
	// The time that the quote was accepted. Measured in seconds since Unix epoch.
	AcceptedAt int `json:"accepted_at,omitempty"`
	// The time that the quote was canceled. Measured in seconds since Unix epoch.
	CanceledAt int `json:"canceled_at,omitempty"`
	// The time that the quote was finalized. Measured in seconds since Unix epoch.
	FinalizedAt int `json:"finalized_at,omitempty"`
}

// QuotesResourceTotalDetailsResourceBreakdown
type QuotesResourceTotalDetailsResourceBreakdown struct {
	core.Model
	// // The aggregated discounts.
	// Discounts []LineItemsDiscountAmount `json:"discounts,omitempty"`
	// // The aggregated tax amounts by rate.
	// Taxes []LineItemsTaxAmount `json:"taxes,omitempty"`
}

// QuotesResourceTotalDetails
type QuotesResourceTotalDetails struct {
	core.Model
	// This is the sum of all the discounts.
	AmountDiscount int `json:"amount_discount,omitempty"`
	// This is the sum of all the shipping amounts.
	AmountShipping int `json:"amount_shipping,omitempty"`
	// This is the sum of all the tax amounts.
	AmountTax int                                          `json:"amount_tax,omitempty"`
	Breakdown *QuotesResourceTotalDetailsResourceBreakdown `json:"breakdown,omitempty" database:"foreignKey:id"`
}

// QuotesResourceRecurring
type QuotesResourceRecurring struct {
	core.Model
	// Total before any discounts or taxes are applied.
	AmountSubtotal int `json:"amount_subtotal,omitempty"`
	// Total after discounts and taxes are applied.
	AmountTotal int `json:"amount_total,omitempty"`
	// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
	Interval string `json:"interval,omitempty"`
	// The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.
	IntervalCount int                        `json:"interval_count,omitempty"`
	TotalDetails  QuotesResourceTotalDetails `json:"total_details,omitempty" database:"foreignKey:id"`
}

// QuotesResourceComputed
type QuotesResourceComputed struct {
	core.Model
	Recurring QuotesResourceRecurring `json:"recurring,omitempty" database:"foreignKey:id"`
	Upfront   QuotesResourceUpfront   `json:"upfront,omitempty" database:"foreignKey:id"`
}

// QuotesResourceUpfront
type QuotesResourceUpfront struct {
	core.Model
	// Total before any discounts or taxes are applied.
	AmountSubtotal int `json:"amount_subtotal,omitempty"`
	// Total after discounts and taxes are applied.
	AmountTotal int `json:"amount_total,omitempty"`
	// LineItems    *Item                      `json:"line_items,omitempty"`
	TotalDetails QuotesResourceTotalDetails `json:"total_details,omitempty" database:"foreignKey:id"`
}

// Quote A Quote is a way to model prices that you'd like to provide to a customer. Once accepted, it will automatically create an invoice, subscription or subscription schedule.
type Quote struct {
	core.Model

	// Total before any discounts or taxes are applied.
	AmountSubtotal int `json:"amount_subtotal,omitempty"`
	// Total after discounts and taxes are applied.
	AmountTotal  int                        `json:"amount_total,omitempty"`
	AutomaticTax QuotesResourceAutomaticTax `json:"automatic_tax,omitempty" database:"foreignKey:id"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay invoices at the end of the subscription cycle or on finalization using the default payment method attached to the subscription or customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically`.
	CollectionMethod string                 `json:"collection_method,omitempty"`
	Computed         QuotesResourceComputed `json:"computed,omitempty" database:"foreignKey:id"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	// A description that will be displayed on the quote PDF.
	Description string `json:"description,omitempty"`
	// The date on which the quote will be canceled if in `open` or `draft` status. Measured in seconds since the Unix epoch.
	ExpiresAt int `json:"expires_at,omitempty"`
	// A footer that will be displayed on the quote PDF.
	Footer string `json:"footer,omitempty"`
	// A header that will be displayed on the quote PDF.
	Header string `json:"header,omitempty"`
	// Unique identifier for the object.
	InvoiceSettings *settings.InvoiceSettingQuoteSetting `json:"invoice_settings,omitempty" database:"foreignKey:id"`
	// LineItems *Item `json:"line_items,omitempty"`
	// A unique number that identifies this particular quote. This number is assigned once the quote is [finalized](https://stripe.com/docs/quotes/overview#finalize).
	Number string `json:"number,omitempty"`
	// The status of the quote.
	Status            string                                          `json:"status,omitempty"`
	StatusTransitions *QuotesResourceStatusTransitions                `json:"status_transitions,omitempty" database:"foreignKey:id"`
	SubscriptionData  *QuotesResourceSubscriptionDataSubscriptionData `json:"subscription_data,omitempty" database:"foreignKey:id"`
	TotalDetails      *QuotesResourceTotalDetails                     `json:"total_details,omitempty" database:"foreignKey:id"`

	CustomerId uuid.UUID          `json:"customer_id" database:"default:null"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
	// DefaultTaxRates tax.TaxRate `json:"default_tax_rates,omitempty" database:"foreignKey:id"`
	// Discounts []Discount `json:"discounts,omitempty" database:"foreignKey:id"`
	InvoiceId              uuid.UUID                                  `json:"invoice_id,omitempty"`
	Invoice                *invoice.Invoice                           `json:"invoice,omitempty" database:"foreignKey:invoice_id"`
	SubscriptionId         uuid.UUID                                  `json:"subscription_id,omitempty"`
	Subscription           *subscription.Subscription                 `json:"subscription,omitempty" database:"foreignKey:subscription_id"`
	SubscriptionScheduleId uuid.UUID                                  `json:"subscription_schedule_id,omitempty"`
	SubscriptionSchedule   *subscriptionSchedule.SubscriptionSchedule `json:"subscription_schedule,omitempty" database:"foreignKey:subscription_schedule_id"`
}
