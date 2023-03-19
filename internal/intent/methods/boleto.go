package methods

import (
	"github.com/driver005/gateway/payment/method/options"
	"github.com/driver005/gateway/utils/paymentFlow"
)

// PaymentIntentPaymentMethodOptionsBoleto struct for PaymentIntentPaymentMethodOptionsBoleto
type PaymentIntentPaymentMethodOptionsBoleto struct {
	options.PaymentMethodOptionsBoleto

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
