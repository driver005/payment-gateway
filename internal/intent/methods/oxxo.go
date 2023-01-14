package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsOxxo struct for PaymentIntentPaymentMethodOptionsOxxo
type PaymentIntentPaymentMethodOptionsOxxo struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsOxxo                            *options.PaymentMethodOptionsOxxo
}
