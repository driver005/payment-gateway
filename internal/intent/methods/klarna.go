package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsKlarna struct for PaymentIntentPaymentMethodOptionsKlarna
type PaymentIntentPaymentMethodOptionsKlarna struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsKlarna                          *options.PaymentMethodOptionsKlarna
}
