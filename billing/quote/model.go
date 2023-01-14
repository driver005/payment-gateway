package quote

import (
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/invoice/settings"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/models"
)

// QuotesResourceListLineItems
type QuotesResourceListLineItems struct {
	// Details about each object.
	// Data []Item `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	// String representing the object's type. Objects of the same type share the same value. Always has the value `list`.
	Object string `json:"object"`
	// The URL where this list can be accessed.
	Url string `json:"url"`
}

// QuotesResourceAutomaticTax
type QuotesResourceAutomaticTax struct {
	// Automatically calculate taxes
	Enabled bool `json:"enabled"`
	// The status of the most recent automated tax calculation for this quote.
	Status string `json:"status,omitempty"`
}

// QuotesResourceSubscriptionDataSubscriptionData
type QuotesResourceSubscriptionDataSubscriptionData struct {
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description,omitempty"`
	// When creating a new subscription, the date of which the subscription schedule will start after the quote is accepted. This date is ignored if it is in the past when the quote is accepted. Measured in seconds since the Unix epoch.
	EffectiveDate int `json:"effective_date,omitempty"`
	// Integer representing the number of trial period days before the customer is charged for the first time.
	TrialPeriodDays int `json:"trial_period_days,omitempty"`
}

// QuotesResourceStatusTransitions
type QuotesResourceStatusTransitions struct {
	// The time that the quote was accepted. Measured in seconds since Unix epoch.
	AcceptedAt int `json:"accepted_at,omitempty"`
	// The time that the quote was canceled. Measured in seconds since Unix epoch.
	CanceledAt int `json:"canceled_at,omitempty"`
	// The time that the quote was finalized. Measured in seconds since Unix epoch.
	FinalizedAt int `json:"finalized_at,omitempty"`
}

// QuotesResourceTotalDetailsResourceBreakdown
type QuotesResourceTotalDetailsResourceBreakdown struct {
	// // The aggregated discounts.
	// Discounts []LineItemsDiscountAmount `json:"discounts"`
	// // The aggregated tax amounts by rate.
	// Taxes []LineItemsTaxAmount `json:"taxes"`
}

// QuotesResourceTotalDetails
type QuotesResourceTotalDetails struct {
	// This is the sum of all the discounts.
	AmountDiscount int32 `json:"amount_discount"`
	// This is the sum of all the shipping amounts.
	AmountShipping int `json:"amount_shipping,omitempty"`
	// This is the sum of all the tax amounts.
	AmountTax int32 `json:"amount_tax"`
	Breakdown *QuotesResourceTotalDetailsResourceBreakdown `json:"breakdown,omitempty"`
}

// QuotesResourceRecurring
type QuotesResourceRecurring struct {
	// Total before any discounts or taxes are applied.
	AmountSubtotal int32 `json:"amount_subtotal"`
	// Total after discounts and taxes are applied.
	AmountTotal int32 `json:"amount_total"`
	// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
	Interval string `json:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.
	IntervalCount int32 `json:"interval_count"`
	TotalDetails QuotesResourceTotalDetails `json:"total_details"`
}

// QuotesResourceComputed
type QuotesResourceComputed struct {
	Recurring  QuotesResourceRecurring `json:"recurring,omitempty"`
	Upfront QuotesResourceUpfront `json:"upfront"`
}

// QuotesResourceUpfront
type QuotesResourceUpfront struct {
	// Total before any discounts or taxes are applied.
	AmountSubtotal int32 `json:"amount_subtotal"`
	// Total after discounts and taxes are applied.
	AmountTotal int32 `json:"amount_total"`
	LineItems *QuotesResourceListLineItems `json:"line_items,omitempty"`
	TotalDetails QuotesResourceTotalDetails `json:"total_details"`
}

// QuotesResourceFromQuote
type QuotesResourceFromQuote struct {
	// Whether this quote is a revision of a different quote.
	IsRevision bool `json:"is_revision"`
	Quote *Quote `json:"quote"`
}


// Quote A Quote is a way to model prices that you'd like to provide to a customer. Once accepted, it will automatically create an invoice, subscription or subscription schedule.
type Quote struct {
	// Total before any discounts or taxes are applied.
	AmountSubtotal int `json:"amount_subtotal"`
	// Total after discounts and taxes are applied.
	AmountTotal int `json:"amount_total"`
	AutomaticTax QuotesResourceAutomaticTax `json:"automatic_tax"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay invoices at the end of the subscription cycle or on finalization using the default payment method attached to the subscription or customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically`.
	CollectionMethod string `json:"collection_method"`
	Computed QuotesResourceComputed `json:"computed"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	Customer customer.Customer `json:"customer,omitempty"`
	// The tax rates applied to this quote.
	DefaultTaxRates []models.TaxRate `json:"default_tax_rates,omitempty"`
	// A description that will be displayed on the quote PDF.
	Description string `json:"description,omitempty"`
	// The discounts applied to this quote.
	
	// Discounts []Discount `json:"discounts"`
	
	// The date on which the quote will be canceled if in `open` or `draft` status. Measured in seconds since the Unix epoch.
	ExpiresAt int `json:"expires_at"`
	// A footer that will be displayed on the quote PDF.
	Footer string `json:"footer,omitempty"`
	FromQuote  QuotesResourceFromQuote `json:"from_quote,omitempty"`
	// A header that will be displayed on the quote PDF.
	Header string `json:"header,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id"`
	Invoice  invoice.Invoice `json:"invoice,omitempty"`
	InvoiceSettings  settings.InvoiceSettingQuoteSetting `json:"invoice_settings,omitempty"`
	LineItems *QuotesResourceListLineItems `json:"line_items,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A unique number that identifies this particular quote. This number is assigned once the quote is [finalized](https://stripe.com/docs/quotes/overview#finalize).
	Number string `json:"number,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The status of the quote.
	Status string `json:"status"`
	StatusTransitions QuotesResourceStatusTransitions `json:"status_transitions"`
	Subscription  *models.Subscription `json:"subscription,omitempty"`
	SubscriptionData QuotesResourceSubscriptionDataSubscriptionData `json:"subscription_data"`
	SubscriptionSchedule  *models.Subscription `json:"subscription_schedule,omitempty"`
	otalDetails QuotesResourceTotalDetails `json:"total_details"`
}
