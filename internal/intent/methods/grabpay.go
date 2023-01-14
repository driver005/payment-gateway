package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsGrabpay struct for PaymentIntentPaymentMethodOptionsGrabpay
type PaymentIntentPaymentMethodOptionsGrabpay struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsGrabpay                         *options.PaymentMethodOptionsGrabpay
}
