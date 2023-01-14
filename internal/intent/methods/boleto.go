package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsBoleto struct for PaymentIntentPaymentMethodOptionsBoleto
type PaymentIntentPaymentMethodOptionsBoleto struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsBoleto                          *options.PaymentMethodOptionsBoleto
}
