package methods

// MandateOptionsParam struct for MandateOptionsParam
type MandateOptionsParam struct {
	CustomMandateUrl    string   `json:"custom_mandate_url,omitempty"`
	DefaultFor          []string `json:"default_for,omitempty"`
	IntervalDescription string   `json:"interval_description,omitempty"`
	PaymentSchedule     string   `json:"payment_schedule,omitempty"`
	TransactionType     string   `json:"transaction_type,omitempty"`
}
