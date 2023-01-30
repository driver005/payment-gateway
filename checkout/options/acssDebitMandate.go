package options

import "github.com/driver005/gateway/core"

// CheckoutAcssDebitMandateOptions
type CheckoutAcssDebitMandateOptions struct {
	core.Model

	// A URL for custom mandate text
	CustomMandateUrl string `json:"custom_mandate_url,omitempty"`
	// List of Stripe products where this mandate can be selected automatically. Returned when the Session is in `setup` mode.
	DefaultFor []string `json:"default_for,omitempty"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description,omitempty"`
	// Payment schedule for the mandate.
	PaymentSchedule string `json:"payment_schedule,omitempty"`
	// Transaction type of the mandate.
	TransactionType string `json:"transaction_type,omitempty"`
}
