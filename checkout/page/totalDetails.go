package page

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/utils/tax"
)

// PaymentPagesCheckoutSessionTotalDetails
type PaymentPagesCheckoutSessionTotalDetails struct {
	core.Model

	// This is the sum of all the discounts.
	AmountDiscount int `json:"amount_discount,omitempty"`
	// This is the sum of all the shipping amounts.
	AmountShipping int `json:"amount_shipping,omitempty"`
	// This is the sum of all the tax amounts.
	AmountTax int `json:"amount_tax,omitempty"`
	// The aggregated discounts.
	Discounts []discount.Discount `json:"discounts,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	// The aggregated tax amounts by rate.
	Taxes []tax.TaxRate `json:"taxes,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
}
