package attempt

import "github.com/driver005/gateway/payment/method"

// SetupAttemptPaymentMethodDetailsCardPresent
type SetupAttemptPaymentMethodDetailsCardPresent struct {
	GeneratedCard method.PaymentMethod `json:"generated_card,omitempty"`
}
