package methods

// MandateSingleUse
type MandateSingleUse struct {
	// On a single use mandate, the amount of the payment.
	Amount int `json:"amount"`
	// On a single use mandate, the currency of the payment.
	Currency string `json:"currency"`
}
