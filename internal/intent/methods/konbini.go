package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsKonbini struct for PaymentIntentPaymentMethodOptionsKonbini
type PaymentIntentPaymentMethodOptionsKonbini struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsKonbini                         *options.PaymentMethodOptionsKonbini
}
