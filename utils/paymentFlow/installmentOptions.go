package paymentFlow

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/payment/method/details"
)

// PaymentFlowsInstallmentOptions
type PaymentFlowsInstallmentOptions struct {
	core.Model

	Enabled bool                                          `json:"enabled,omitempty"`
	Plan    *details.PaymentMethodDetailsCardInstallments `json:"plan,omitempty" database:"foreignKey:id"`
}
