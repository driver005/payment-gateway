package methods

import "github.com/driver005/gateway/utils/paymentFlow"

// PaymentIntentPaymentMethodOptionsAcssDebit1 struct for PaymentIntentPaymentMethodOptionsAcssDebit1
type PaymentIntentPaymentMethodOptionsAcssDebit1 struct {
	PaymentIntentPaymentMethodOptionsAcssDebit

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
