package methods

import "github.com/driver005/gateway/utils/paymentFlow"

// PaymentIntentPaymentMethodOptionsLink1 struct for PaymentIntentPaymentMethodOptionsLink1
type PaymentIntentPaymentMethodOptionsLink1 struct {
	PaymentIntentPaymentMethodOptionsLink

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
