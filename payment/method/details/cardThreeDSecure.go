package details

import "github.com/driver005/gateway/models"

// PaymentMethodDetailsCardThreeDSecure Populated if this transaction used 3D Secure authentication.
type PaymentMethodDetailsCardThreeDSecure struct {
	ThreeDSecureDetails *models.ThreeDSecureDetails
}
