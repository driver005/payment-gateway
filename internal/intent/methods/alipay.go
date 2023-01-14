package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsAlipay struct for PaymentIntentPaymentMethodOptionsAlipay
type PaymentIntentPaymentMethodOptionsAlipay struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsAlipay                          *options.PaymentMethodOptionsAlipay
}
