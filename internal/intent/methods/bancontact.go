package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsBancontact struct for PaymentIntentPaymentMethodOptionsBancontact
type PaymentIntentPaymentMethodOptionsBancontact struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsBancontact                      *options.PaymentMethodOptionsBancontact
}
