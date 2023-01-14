package methods

import "github.com/driver005/gateway/models"

// PaymentMethodSepaDebitGeneratedFrom Information about the object that generated this PaymentMethod.
type PaymentMethodSepaDebitGeneratedFrom struct {
	Charge       *models.Charge                             `json:"charge,omitempty"`
	SetupAttempt PaymentMethodCardGeneratedCardSetupAttempt `json:"setup_attempt,omitempty"`
}
