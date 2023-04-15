package page

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/shipping"
	"github.com/driver005/gateway/utils/tax"
)

// PaymentPagesCheckoutSessionShippingCost
type PaymentPagesCheckoutSessionShippingCost struct {
	core.Model

	// Total shipping cost before any discounts or taxes are applied.
	AmountSubtotal int `json:"amount_subtota"`
	// Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0.
	AmountTax int `json:"amount_tax"`
	// Total shipping cost after discounts and taxes are applied.
	AmountTotal  int                    `json:"amount_tota"`
	ShippingRate *shipping.ShippingRate `json:"shipping_rate,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// The taxes applied to the shipping rate.
	Taxes []tax.TaxRate `json:"taxes,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
}
