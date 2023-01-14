package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsWechatPay struct for PaymentIntentPaymentMethodOptionsWechatPay
type PaymentIntentPaymentMethodOptionsWechatPay struct {
	PaymentIntentTypeSpecificPaymentMethodOptionsClient *PaymentIntentTypeSpecificPaymentMethodOptionsClient
	PaymentMethodOptionsWechatPay                       *options.PaymentMethodOptionsWechatPay
}
