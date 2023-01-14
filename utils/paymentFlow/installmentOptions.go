package paymentFlow

import "github.com/driver005/gateway/payment/method/details"

// PaymentFlowsInstallmentOptions
type PaymentFlowsInstallmentOptions struct {
	Enabled bool                                              `json:"enabled"`
	Plan    *details.PaymentMethodDetailsCardInstallmentsPlan `json:"plan,omitempty"`
}
