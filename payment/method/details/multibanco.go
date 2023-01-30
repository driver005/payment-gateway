package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsMultibanco
type PaymentMethodDetailsMultibanco struct {
	core.Model

	// Entity number associated with this Multibanco payment.
	Entity string `json:"entity,omitempty"`
	// Reference number associated with this Multibanco payment.
	Reference string `json:"reference,omitempty"`
}
