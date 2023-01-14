package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsFpx struct for PaymentIntentPaymentMethodOptionsFpx
type PaymentIntentPaymentMethodOptionsFpx struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsFpx                             *options.PaymentMethodOptionsFpx
}
