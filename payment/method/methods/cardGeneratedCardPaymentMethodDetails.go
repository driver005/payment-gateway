package methods

import "github.com/driver005/gateway/payment/method/details"

// PaymentMethodCardGeneratedCardPaymentMethodDetails Transaction-specific details of the payment method used in the payment.
type PaymentMethodCardGeneratedCardPaymentMethodDetails struct {
	CardPresent *details.PaymentMethodDetailsCardPresent `json:"card_present,omitempty"`
	// The type of payment method transaction-specific details from the transaction that generated this `card` payment method. Always `card_present`.
	Type string `json:"type"`
}
