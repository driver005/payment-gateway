package paymentFlow

import "github.com/driver005/gateway/core"

// PaymentFlowsAutomaticPaymentMethodsPaymentIntent
type PaymentFlowsAutomaticPaymentMethodsPaymentIntent struct {
	core.Model

	// Automatically calculates compatible payment methods
	Enabled bool `json:"enabled,omitempty"`
}
