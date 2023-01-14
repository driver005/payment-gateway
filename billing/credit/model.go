package credit

import (
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/refund"
	"github.com/driver005/gateway/models"
)

// CreditNoteTaxAmount
type CreditNoteTaxAmount struct {
	// The amount, in %s, of the tax.
	Amount int `json:"amount"`
	// Whether this tax amount is inclusive or exclusive.
	Inclusive bool `json:"inclusive"`
	TaxRate models.TaxRate `json:"tax_rate"`
}

// CreditNoteLineItemParams struct for CreditNoteLineItemParams
type CreditNoteLineItemParams struct {
	Amount *int `json:"amount,omitempty"`
	Description *string `json:"description,omitempty"`
	InvoiceLineItem *string `json:"invoice_line_item,omitempty"`
	Quantity *int `json:"quantity,omitempty"`
	TaxRates *string `json:"tax_rates,omitempty"`
	Type string `json:"type"`
	UnitAmount *int `json:"unit_amount,omitempty"`
	UnitAmountDecimal *float64 `json:"unit_amount_decimal,omitempty"`
}

// CreditNoteLineItem
type CreditNoteLineItem struct {
	// The integer amount in %s representing the gross amount being credited for this line item, excluding (exclusive) tax and discounts.
	Amount int `json:"amount"`
	// The integer amount in %s representing the amount being credited for this line item, excluding all tax and discounts.
	AmountExcludingTax int `json:"amount_excluding_tax,omitempty"`
	// Description of the item being credited.
	Description string `json:"description,omitempty"`
	// The integer amount in %s representing the discount being credited for this line item.
	DiscountAmount int `json:"discount_amount"`
	// The amount of discount calculated per discount for this line item
	// DiscountAmounts []DiscountsResourceDiscountAmount `json:"discount_amounts"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// ID of the invoice line item being credited
	InvoiceLineItem *string `json:"invoice_line_item,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The number of units of product being credited.
	Quantity int `json:"quantity,omitempty"`
	// The amount of tax calculated per tax rate for this line item
	TaxAmounts []CreditNoteTaxAmount `json:"tax_amounts"`
	// The tax rates which apply to the line item.
	TaxRates []models.TaxRate `json:"tax_rates"`
	// The type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. When the type is `invoice_line_item` there is an additional `invoice_line_item` property on the resource the value of which is the id of the credited line item on the invoice.
	Type string `json:"type"`
	// The cost of each unit of product being credited.
	UnitAmount int `json:"unit_amount,omitempty"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,omitempty"`
	// The amount in %s representing the unit amount being credited for this line item, excluding all tax and discounts.
	UnitAmountExcludingTax float64 `json:"unit_amount_excluding_tax,omitempty"`
}

// CreditNote Issue a credit note to adjust an invoice's amount after the invoice is finalized.  Related guide: [Credit Notes](https://stripe.com/docs/billing/invoices/credit-notes).
type CreditNote struct {
	// The integer amount in %s representing the total amount of the credit note, including tax.
	Amount int `json:"amount"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency"`
	Customer customer.Customer `json:"customer"`
	CustomerBalanceTransaction balance.BalanceTransaction `json:"customer_balance_transaction,omitempty"`
	// The integer amount in %s representing the total amount of discount that was credited.
	DiscountAmount int `json:"discount_amount"`
	// The aggregate amounts calculated per discount for all line items.
	//TODO: add Discount
	// DiscountAmounts []Discount `json:"discount_amounts"`
	// Unique identifier for the object.
	Id string `json:"id"`
	Invoice invoice.Invoice `json:"invoice"`
	Lines []CreditNoteLineItem `json:"lines"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Customer-facing text that appears on the credit note PDF.
	Memo string `json:"memo,omitempty"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// A unique number that identifies this particular credit note and appears on the PDF of the credit note and its associated invoice.
	Number string `json:"number"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Amount that was credited outside of Stripe.
	OutOfBandAmount int `json:"out_of_band_amount,omitempty"`
	// The link to download the PDF of the credit note.
	Pdf string `json:"pdf"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason string `json:"reason,omitempty"`
	Refund refund.Refund `json:"refund,omitempty"`
	// Status of this credit note, one of `issued` or `void`. Learn more about [voiding credit notes](https://stripe.com/docs/billing/invoices/credit-notes#voiding).
	Status string `json:"status"`
	// The integer amount in %s representing the amount of the credit note, excluding exclusive tax and invoice level discounts.
	Subtotal int `json:"subtotal"`
	// The integer amount in %s representing the amount of the credit note, excluding all tax and invoice level discounts.
	SubtotalExcludingTax int `json:"subtotal_excluding_tax,omitempty"`
	// The aggregate amounts calculated per tax rate for all line items.
	TaxAmounts []CreditNoteTaxAmount `json:"tax_amounts"`
	// The integer amount in %s representing the total amount of the credit note, including tax and all discount.
	Total int `json:"total"`
	// The integer amount in %s representing the total amount of the credit note, excluding tax, but including discounts.
	TotalExcludingTax int `json:"total_excluding_tax,omitempty"`
	// Type of this credit note, one of `pre_payment` or `post_payment`. A `pre_payment` credit note means it was issued when the invoice was open. A `post_payment` credit note means it was issued when the invoice was paid.
	Type string `json:"type"`
	// The time that the credit note was voided.
	VoidedAt int `json:"voided_at,omitempty"`
}
