package methods

import "github.com/driver005/gateway/payment/method/options"

// PaymentIntentPaymentMethodOptionsCardInstallments Installment details for this payment (Mexico only).  For more information, see the [installments integration guide](https://stripe.com/docs/payments/installments).
type PaymentIntentPaymentMethodOptionsCardInstallments struct {
	PaymentMethodOptionsCardInstallments *options.PaymentMethodOptionsCardInstallments
}
