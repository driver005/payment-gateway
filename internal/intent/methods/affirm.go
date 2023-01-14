package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsAffirm struct for PaymentIntentPaymentMethodOptionsAffirm
type PaymentIntentPaymentMethodOptionsAffirm struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsAffirm                          *options.PaymentMethodOptionsAffirm
}
