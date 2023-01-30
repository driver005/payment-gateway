package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/payment/method/details"
)

// SetupAttemptPaymentMethodDetailsCard
type SetupAttemptPaymentMethodDetailsCard struct {
	core.Model

	ThreeDSecure details.ThreeDSecureDetails `json:"three_d_secure,omitempty" database:"foreignKey:id"`
}
