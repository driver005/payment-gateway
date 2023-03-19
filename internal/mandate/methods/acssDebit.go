package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// MandateAcssDebit
type MandateAcssDebit struct {
	core.Model

	// List of Stripe products where this mandate can be selected automatically.
	DefaultFor pq.StringArray `json:"default_for,omitempty" database:"type:varchar(64)[]"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description,omitempty"`
	// Payment schedule for the mandate.
	PaymentSchedule string `json:"payment_schedule,omitempty"`
	// Transaction type of the mandate.
	TransactionType string `json:"transaction_type,omitempty"`
}
