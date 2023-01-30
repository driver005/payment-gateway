package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/paymentFlow"
)

// PaymentIntentPaymentMethodOptionsInteracPresent struct for PaymentIntentPaymentMethodOptionsInteracPresent
type PaymentIntentPaymentMethodOptionsInteracPresent struct {
	core.Model

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string                                      `json:"capture_method,omitempty"`
	Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`

	SetupFutureUsage string `json:"setup_future_usage,omitempty"`
}
