package methods

import "github.com/driver005/gateway/utils/paymentFlow"

// PaymentIntentPaymentMethodOptionsUsBankAccount1 struct for PaymentIntentPaymentMethodOptionsUsBankAccount1
type PaymentIntentPaymentMethodOptionsUsBankAccount1 struct {
	PaymentIntentPaymentMethodOptionsUsBankAccount

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
