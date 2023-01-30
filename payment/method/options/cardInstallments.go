package options

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/payment/method/details"
)

// PaymentMethodOptionsCardInstallments
type PaymentMethodOptionsCardInstallments struct {
	core.Model

	// Installment plans that may be selected for this PaymentIntent.
	AvailablePlans []details.PaymentMethodDetailsCardInstallments `json:"available_plans,omitempty" database:"foreignKey:id"`
	// Whether Installments are enabled for this PaymentIntent.
	Enabled bool                                          `json:"enabled,omitempty"`
	Plan    *details.PaymentMethodDetailsCardInstallments `json:"plan,omitempty" database:"foreignKey:id"`
}
