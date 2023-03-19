package options

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/shipping"
)

// PaymentLinksResourceShippingOption
type PaymentLinksResourceShippingOption struct {
	core.Model

	// A non-negative integer in cents representing how much to charge.
	ShippingAmount int                    `json:"shipping_amount,omitempty"`
	ShippingRate   *shipping.ShippingRate `json:"shipping_rate,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
