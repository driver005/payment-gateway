package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsP24 struct for PaymentIntentPaymentMethodOptionsP24
type PaymentIntentPaymentMethodOptionsP24 struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsP24                             *options.PaymentMethodOptionsP24
}
