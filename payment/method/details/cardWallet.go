package details

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/address"
)

type PaymentMethodCardWalletDetails struct {
	core.Model

	BillingAddress  *address.Address `json:"billing_address,omitempty" database:"foreignKey:id"`
	Email           string           `json:"email,omitempty"`
	Name            string           `json:"name,omitempty"`
	ShippingAddress *address.Address `json:"shipping_address,omitempty" database:"foreignKey:id"`
}

type PaymentMethodDetailsCardWallet struct {
	core.Model

	AmexExpressCheckout *PaymentMethodCardWalletDetails `json:"amex_express_checkout,omitempty" database:"foreignKey:id"`
	ApplePay            *PaymentMethodCardWalletDetails `json:"apple_pay,omitempty" database:"foreignKey:id"`
	DynamicLast4        string                          `json:"dynamic_last4,omitempty"`
	GooglePay           *PaymentMethodCardWalletDetails `json:"google_pay,omitempty" database:"foreignKey:id"`
	Masterpass          *PaymentMethodCardWalletDetails `json:"masterpass,omitempty" database:"foreignKey:id"`
	SamsungPay          *PaymentMethodCardWalletDetails `json:"samsung_pay,omitempty" database:"foreignKey:id"`
	Type                string                          `json:"type,omitempty"`
	VisaCheckout        *PaymentMethodCardWalletDetails `json:"visa_checkout,omitempty" database:"foreignKey:id"`
}
