package methods

// PaymentMethodKlarnaDob The customer's date of birth, if provided.
type PaymentMethodKlarnaDob struct {
	// The day of birth, between 1 and 31.
	Day int `json:"day,omitempty"`
	// The month of birth, between 1 and 12.
	Month int `json:"month,omitempty"`
	// The four-digit year of birth.
	Year int `json:"year,omitempty"`
}
