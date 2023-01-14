package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsPromptpay struct for PaymentIntentPaymentMethodOptionsPromptpay
type PaymentIntentPaymentMethodOptionsPromptpay struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsPromptpay                       *options.PaymentMethodOptionsPromptpay
}
