package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsBacsDebit struct for PaymentIntentPaymentMethodOptionsBacsDebit
type PaymentIntentPaymentMethodOptionsBacsDebit struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsBacsDebit                       *options.PaymentMethodOptionsBacsDebit
}
