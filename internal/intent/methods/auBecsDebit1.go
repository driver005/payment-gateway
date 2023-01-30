package methods

import "github.com/driver005/gateway/utils/paymentFlow"

// PaymentIntentPaymentMethodOptionsAuBecsDebit1 struct for PaymentIntentPaymentMethodOptionsAuBecsDebit1
type PaymentIntentPaymentMethodOptionsAuBecsDebit1 struct {
	PaymentIntentPaymentMethodOptionsAuBecsDebit

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
