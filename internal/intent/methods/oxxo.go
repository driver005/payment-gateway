package methods

import (
	"github.com/driver005/gateway/payment/method/options"
	"github.com/driver005/gateway/utils/paymentFlow"
)

// PaymentIntentPaymentMethodOptionsOxxo struct for PaymentIntentPaymentMethodOptionsOxxo
type PaymentIntentPaymentMethodOptionsOxxo struct {
	options.PaymentMethodOptionsOxxo

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
