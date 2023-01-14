package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsSofort struct for PaymentIntentPaymentMethodOptionsSofort
type PaymentIntentPaymentMethodOptionsSofort struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsSofort                          *options.PaymentMethodOptionsSofort
}
