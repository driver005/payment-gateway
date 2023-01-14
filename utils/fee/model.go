package fee

// Fee
type Fee struct {
	// Amount of the fee, in cents.
	Amount int `json:"amount"`
	// ID of the Connect application that earned the fee.
	Application string `json:"application,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// Type of the fee, one of: `application_fee`, `stripe_fee` or `tax`.
	Type string `json:"type"`
}
