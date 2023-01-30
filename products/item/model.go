package item

import (
	"github.com/driver005/gateway/billing/invoice/invoices"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/utils/tax"
)

type LineItem struct {
	core.Model

	Amount                 int                                      `json:"amount"`
	AmountExcludingTax     int                                      `json:"amount_excluding_tax,omitempty"`
	Currency               string                                   `json:"currency"`
	Description            string                                   `json:"description,omitempty"`
	Discountable           bool                                     `json:"discountable"`
	Discounts              []discount.Discount                      `json:"discounts,omitempty"`
	InvoiceItem            string                                   `json:"invoice_item,omitempty"`
	Price                  *price.Price                             `json:"price,omitempty"`
	Proration              bool                                     `json:"proration"`
	ProrationDetails       *invoices.InvoicesLineItemsCreditedItems `json:"proration_details,omitempty"`
	Quantity               int                                      `json:"quantity,omitempty"`
	Subscription           string                                   `json:"subscription,omitempty"`
	SubscriptionItem       string                                   `json:"subscription_item,omitempty"`
	TaxRates               []tax.TaxRate                            `json:"tax_rates,omitempty"`
	Type                   string                                   `json:"type"`
	UnitAmountExcludingTax float64                                  `json:"unit_amount_excluding_tax,omitempty"`
}
