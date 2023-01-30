package methods

import "github.com/driver005/gateway/core"

// MandateOptionsParam struct for MandateOptionsParam
type MandateOptionsParam struct {
	core.Model

	CustomMandateUrl    string   `json:"custom_mandate_url,omitempty"`
	DefaultFor          []string `json:"default_for,omitempty" database:"type:text[]"`
	IntervalDescription string   `json:"interval_description,omitempty"`
	PaymentSchedule     string   `json:"payment_schedule,omitempty"`
	TransactionType     string   `json:"transaction_type,omitempty"`
}
