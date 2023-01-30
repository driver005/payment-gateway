package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionCardAwaitNotification
type PaymentIntentNextActionCardAwaitNotification struct {
	core.Model

	// The time that payment will be attempted. If customer approval is required, they need to provide approval before this time.
	ChargeAttemptAt int `json:"charge_attempt_at,omitempty"`
	// For payments greater than INR 15000, the customer must provide explicit approval of the payment with their bank. For payments of lower amount, no customer action is required.
	CustomerApprovalRequired bool `json:"customer_approval_required,omitempty"`
}
