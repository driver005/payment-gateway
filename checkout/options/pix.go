package checkout

// CheckoutPixPaymentMethodOptions
type CheckoutPixPaymentMethodOptions struct {
	// The number of seconds after which Pix payment will expire.
	ExpiresAfterSeconds int `json:"expires_after_seconds,omitempty"`
}
