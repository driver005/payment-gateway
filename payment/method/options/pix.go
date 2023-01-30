package options

import "github.com/driver005/gateway/core"

// PaymentMethodOptionsPix
type PaymentMethodOptionsPix struct {
	core.Model

	// The number of seconds (between 10 and 1209600) after which Pix payment will expire.
	ExpiresAfterSeconds int `json:"expires_after_seconds,omitempty"`
	// The timestamp at which the Pix expires.
	ExpiresAt int `json:"expires_at,omitempty"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.  Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.  When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage string `json:"setup_future_usage,omitempty"`
}
