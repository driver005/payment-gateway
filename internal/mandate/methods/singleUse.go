package methods

import "github.com/driver005/gateway/core"

// MandateSingleUse
type MandateSingleUse struct {
	core.Model

	// On a single use mandate, the amount of the payment.
	Amount int `json:"amount,omitempty"`
	// On a single use mandate, the currency of the payment.
	Currency string `json:"currency,omitempty"`
}
