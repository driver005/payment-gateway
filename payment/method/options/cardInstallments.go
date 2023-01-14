package options

import "github.com/driver005/gateway/payment/method/details"

// PaymentMethodOptionsCardInstallments
type PaymentMethodOptionsCardInstallments struct {
	// Installment plans that may be selected for this PaymentIntent.
	AvailablePlans []details.PaymentMethodDetailsCardInstallmentsPlan `json:"available_plans,omitempty"`
	// Whether Installments are enabled for this PaymentIntent.
	Enabled bool                                     `json:"enabled"`
	Plan    PaymentMethodOptionsCardInstallmentsPlan `json:"plan,omitempty"`
}
