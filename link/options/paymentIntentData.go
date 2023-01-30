package options

import "github.com/driver005/gateway/core"

// PaymentLinksResourcePaymentIntentData
type PaymentLinksResourcePaymentIntentData struct {
	core.Model

	// Indicates when the funds will be captured from the customer's account.
	CaptureMethod string `json:"capture_method,omitempty"`
	// Indicates that you intend to make future payments with the payment method collected during checkout.
	SetupFutureUsage string `json:"setup_future_usage,omitempty"`
}
