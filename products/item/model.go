package item

import (
	"time"

	"github.com/driver005/gateway/billing/invoice/invoices"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/products/product"
	"github.com/driver005/gateway/utils/tax"
	"github.com/google/uuid"
)

type LineItem struct {
	core.Model

	Amount                 int                                      `json:"amount,omitempty"`
	AmountExcludingTax     int                                      `json:"amount_excluding_tax,omitempty"`
	Currency               string                                   `json:"currency,omitempty"`
	Description            string                                   `json:"description,omitempty"`
	Discountable           bool                                     `json:"discountable,omitempty"`
	Discounts              []discount.Discount                      `json:"discounts,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	PeriodStart            time.Time                                `json:"period_start,omitempty"`
	PeriodEnd              time.Time                                `json:"period_end,omitempty"`
	Proration              bool                                     `json:"proration,omitempty"`
	ProrationDetails       *invoices.InvoicesLineItemsCreditedItems `json:"proration_details,omitempty" database:"foreignKey:id"`
	Quantity               int                                      `json:"quantity,omitempty"`
	Subscription           string                                   `json:"subscription,omitempty"`
	SubscriptionItem       string                                   `json:"subscription_item,omitempty"`
	TaxRates               []tax.TaxRate                            `json:"tax_rates,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	Type                   string                                   `json:"type,omitempty"`
	UnitAmountExcludingTax float64                                  `json:"unit_amount_excluding_tax,omitempty"`

	ProductId *uuid.UUID       `json:"product_id,omitempty" swaggerignore:"true"`
	Product   *product.Product `json:"product,omitempty" database:"foreignKey:product_id"`
	PriceId   *uuid.UUID       `json:"price_id,omitempty" swaggerignore:"true"`
	Price     *price.Price     `json:"price,omitempty" database:"foreignKey:price_id"`
}
