package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsPix struct for PaymentIntentPaymentMethodOptionsPix
type PaymentIntentPaymentMethodOptionsPix struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsPix                             *options.PaymentMethodOptionsPix
}
