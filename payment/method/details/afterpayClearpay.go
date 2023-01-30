package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsAfterpayClearpay
type PaymentMethodDetailsAfterpayClearpay struct {
	core.Model

	// Order identifier shown to the merchant in Afterpay’s online portal.
	Reference string `json:"reference,omitempty"`
}
