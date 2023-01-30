package paymentFlow

import "github.com/driver005/gateway/core"

// PaymentFlowsAmountDetails
type PaymentFlowsAmountDetails struct {
	core.Model

	TipAmount int `json:"tip_amount,omitempty"`
}
