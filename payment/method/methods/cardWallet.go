package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/address"
)

type PaymentMethodCardWalletDetails struct {
	core.Model

	BillingAddress  *address.Address `json:"billing_address,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Email           string           `json:"email,omitempty"`
	Name            string           `json:"name,omitempty"`
	ShippingAddress *address.Address `json:"shipping_address,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}

type PaymentMethodCardWallet struct {
	core.Model

	AmexExpressCheckout *PaymentMethodCardWalletDetails `json:"amex_express_checkout,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	ApplePay            *PaymentMethodCardWalletDetails `json:"apple_pay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	DynamicLast4        string                          `json:"dynamic_last4,omitempty"`
	GooglePay           *PaymentMethodCardWalletDetails `json:"google_pay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Masterpass          *PaymentMethodCardWalletDetails `json:"masterpass,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SamsungPay          *PaymentMethodCardWalletDetails `json:"samsung_pay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Type                string                          `json:"type"`
	VisaCheckout        *PaymentMethodCardWalletDetails `json:"visa_checkout,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
