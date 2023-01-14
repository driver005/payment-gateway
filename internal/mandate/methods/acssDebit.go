package methods

// MandateAcssDebit
type MandateAcssDebit struct {
	// List of Stripe products where this mandate can be selected automatically.
	DefaultFor []string `json:"default_for,omitempty"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description,omitempty"`
	// Payment schedule for the mandate.
	PaymentSchedule string `json:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType string `json:"transaction_type"`
}
