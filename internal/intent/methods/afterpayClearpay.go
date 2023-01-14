package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsAfterpayClearpay struct for PaymentIntentPaymentMethodOptionsAfterpayClearpay
type PaymentIntentPaymentMethodOptionsAfterpayClearpay struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsAfterpayClearpay                *options.PaymentMethodOptionsAfterpayClearpay
}
