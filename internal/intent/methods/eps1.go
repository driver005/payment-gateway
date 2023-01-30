package methods

import "github.com/driver005/gateway/utils/paymentFlow"

// PaymentIntentPaymentMethodOptionsEps1 struct for PaymentIntentPaymentMethodOptionsEps1
type PaymentIntentPaymentMethodOptionsEps1 struct {
	PaymentIntentPaymentMethodOptionsEps

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
