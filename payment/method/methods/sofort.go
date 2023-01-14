package methods

// PaymentMethodSofort
type PaymentMethodSofort struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country,omitempty"`
}
