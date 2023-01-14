package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsGiropay struct for PaymentIntentPaymentMethodOptionsGiropay
type PaymentIntentPaymentMethodOptionsGiropay struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsGiropay                         *options.PaymentMethodOptionsGiropay
}
