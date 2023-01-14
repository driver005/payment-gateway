package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsCardPresent struct for PaymentIntentPaymentMethodOptionsCardPresent
type PaymentIntentPaymentMethodOptionsCardPresent struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsCardPresent                     *options.PaymentMethodOptionsCardPresent
}
