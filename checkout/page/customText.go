package page

import "github.com/driver005/gateway/core"

// PaymentPagesCheckoutSessionCustomText
type PaymentPagesCheckoutSessionCustomText struct {
	core.Model

	ShippingAddress string `json:"shipping_address,omitempty"`
	Submit          string `json:"submit,omitempty"`
}
