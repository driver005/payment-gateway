package options

import "github.com/driver005/gateway/payment/method/details"

// PaymentMethodOptionsCardInstallmentsPlan Installment plan selected for this PaymentIntent.
type PaymentMethodOptionsCardInstallmentsPlan struct {
	PaymentMethodDetailsCardInstallmentsPlan *details.PaymentMethodDetailsCardInstallmentsPlan
}
