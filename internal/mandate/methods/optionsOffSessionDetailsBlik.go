package methods

// MandateOptionsOffSessionDetailsBlik
type MandateOptionsOffSessionDetailsBlik struct {
	// Amount of each recurring payment.
	Amount int `json:"amount,omitempty"`
	// Currency of each recurring payment.
	Currency string `json:"currency,omitempty"`
	// Frequency interval of each recurring payment.
	Interval string `json:"interval,omitempty"`
	// Frequency indicator of each recurring payment.
	IntervalCount int `json:"interval_count,omitempty"`
}
