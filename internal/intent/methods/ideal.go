package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsIdeal struct for PaymentIntentPaymentMethodOptionsIdeal
type PaymentIntentPaymentMethodOptionsIdeal struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsIdeal                           *options.PaymentMethodOptionsIdeal
}
