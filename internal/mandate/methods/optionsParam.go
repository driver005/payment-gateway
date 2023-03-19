package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// MandateOptionsParam struct for MandateOptionsParam
type MandateOptionsParam struct {
	core.Model

	CustomMandateUrl    string         `json:"custom_mandate_url,omitempty"`
	DefaultFor          pq.StringArray `json:"default_for,omitempty" database:"type:varchar(64)[]"`
	IntervalDescription string         `json:"interval_description,omitempty"`
	PaymentSchedule     string         `json:"payment_schedule,omitempty"`
	TransactionType     string         `json:"transaction_type,omitempty"`
}
