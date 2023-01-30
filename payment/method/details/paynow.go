package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsPaynow
type PaymentMethodDetailsPaynow struct {
	core.Model

	// Reference number associated with this PayNow payment
	Reference string `json:"reference,omitempty"`
}
