package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsPaynow struct for PaymentIntentPaymentMethodOptionsPaynow
type PaymentIntentPaymentMethodOptionsPaynow struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsPaynow                          *options.PaymentMethodOptionsPaynow
}
