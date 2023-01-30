package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsCardInstallments
type PaymentMethodDetailsCardInstallments struct {
	core.Model

	// For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.
	PlanCount int `json:"plan_count,omitempty"`
	// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card. One of `month`.
	PlanInterval string `json:"plan_interval,omitempty"`
	// Type of installment plan, one of `fixed_count`.
	PlanType string `json:"plan_type,omitempty"`
}
