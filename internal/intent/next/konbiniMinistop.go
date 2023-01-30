package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionKonbiniMinistop
type PaymentIntentNextActionKonbiniMinistop struct {
	core.Model

	// The confirmation number.
	ConfirmationNumber string `json:"confirmation_number,omitempty"`
	// The payment code.
	PaymentCode string `json:"payment_code,omitempty"`
}
