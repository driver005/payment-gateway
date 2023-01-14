package attempt

import "github.com/driver005/gateway/models"

// SetupAttemptPaymentMethodDetailsCard
type SetupAttemptPaymentMethodDetailsCard struct {
	ThreeDSecure models.ThreeDSecureDetails `json:"three_d_secure,omitempty"`
}
