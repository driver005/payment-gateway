package methods

import "github.com/driver005/gateway/models"

// PaymentMethodCardGeneratedCardSetupAttempt The ID of the SetupAttempt that generated this PaymentMethod, if any.
type PaymentMethodCardGeneratedCardSetupAttempt struct {
	SetupAttempt *models.SetupAttempt
	string       string
}
