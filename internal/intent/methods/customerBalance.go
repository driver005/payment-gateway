package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsCustomerBalance struct for PaymentIntentPaymentMethodOptionsCustomerBalance
type PaymentIntentPaymentMethodOptionsCustomerBalance struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsCustomerBalance                 *options.PaymentMethodOptionsCustomerBalance
}
