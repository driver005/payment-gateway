package methods

import "github.com/driver005/gateway/utils/paymentFlow"

// PaymentIntentPaymentMethodOptionsSepaDebit1 struct for PaymentIntentPaymentMethodOptionsSepaDebit1
type PaymentIntentPaymentMethodOptionsSepaDebit1 struct {
	PaymentIntentPaymentMethodOptionsSepaDebit

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
