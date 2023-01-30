package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/payment/method"
)

// SetupAttemptPaymentMethodDetailsCardPresent
type SetupAttemptPaymentMethodDetailsCardPresent struct {
	core.Model

	GeneratedCard method.PaymentMethod `json:"generated_card,omitempty" database:"foreignKey:id"`
}
