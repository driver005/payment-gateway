package methods

import "github.com/driver005/gateway/core"

// PaymentIntentPaymentMethodOptionsMandateOptionsAcssDebit
type PaymentIntentPaymentMethodOptionsMandateOptionsAcssDebit struct {
	core.Model

	// A URL for custom mandate text
	CustomMandateUrl string `json:"custom_mandate_url,omitempty"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description,omitempty"`
	// Payment schedule for the mandate.
	PaymentSchedule string `json:"payment_schedule,omitempty"`
	// Transaction type of the mandate.
	TransactionType string `json:"transaction_type,omitempty"`
}
