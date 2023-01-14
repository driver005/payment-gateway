package methods

import "github.com/driver005/gateway/models"

// PaymentMethodCardThreeDSecureUsage Contains details on how this Card maybe be used for 3D Secure authentication.
type PaymentMethodCardThreeDSecureUsage struct {
	ThreeDSecureUsage *models.ThreeDSecureUsage
}
