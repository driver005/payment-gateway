package options

import "github.com/driver005/gateway/core"

// PaymentLinksResourceCustomText
type PaymentLinksResourceCustomText struct {
	core.Model

	ShippingAddress string `json:"shipping_address,omitempty"`
	Submit          string `json:"submit,omitempty"`
}
