package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsOxxo
type PaymentMethodDetailsOxxo struct {
	core.Model

	// OXXO reference number
	Number string `json:"number,omitempty"`
}
