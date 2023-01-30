package options

import "github.com/driver005/gateway/core"

// CheckoutPixPaymentMethodOptions
type CheckoutPixPaymentMethodOptions struct {
	core.Model

	// The number of seconds after which Pix payment will expire.
	ExpiresAfterSeconds int `json:"expires_after_seconds,omitempty"`
}
