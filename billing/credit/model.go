package credit

import (
	"time"

	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/refund"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/utils/tax"
	"github.com/google/uuid"
)

// CreditNoteLineItem
type CreditNoteLineItem struct {
	core.Model

	// The integer amount in %s representing the gross amount being credited for this line item, excluding (exclusive) tax and discounts.
	Amount int `json:"amount,omitempty"`
	// The integer amount in %s representing the amount being credited for this line item, excluding all tax and discounts.
	AmountExcludingTax int `json:"amount_excluding_tax,omitempty"`
	// Description of the item being credited.
	Description string `json:"description,omitempty"`
	// The integer amount in %s representing the discount being credited for this line item.
	DiscountAmount int `json:"discount_amount,omitempty"`
	// The amount of discount calculated per discount for this line item
	DiscountAmounts []discount.Discount `json:"discount_amounts,omitempty"`
	// ID of the invoice line item being credited
	InvoiceLineItem string `json:"invoice_line_item,omitempty"`
	// The number of units of product being credited.
	Quantity int `json:"quantity,omitempty"`
	// The amount of tax calculated per tax rate for this line item
	TaxAmounts []tax.TaxRate `json:"tax_amounts,omitempty" database:"foreignKey:id"`
	// The tax rates which apply to the line item.
	TaxRates tax.TaxRate `json:"tax_rates,omitempty" database:"foreignKey:id"`
	// The type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. When the type is `invoice_line_item` there is an additional `invoice_line_item` property on the resource the value of which is the id of the credited line item on the invoice.
	Type string `json:"type,omitempty"`
	// The cost of each unit of product being credited.
	UnitAmount int `json:"unit_amount,omitempty"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,omitempty"`
	// The amount in %s representing the unit amount being credited for this line item, excluding all tax and discounts.
	UnitAmountExcludingTax float64 `json:"unit_amount_excluding_tax,omitempty"`
}

// CreditNote Issue a credit note to adjust an invoice's amount after the invoice is finalized.  Related guide: [Credit Notes](https://stripe.com/docs/billing/invoices/credit-notes).
type CreditNote struct {
	core.Model

	// The integer amount in %s representing the total amount of the credit note, including tax.
	Amount int `json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	// The integer amount in %s representing the total amount of discount that was credited.
	DiscountAmount int `json:"discount_amount,omitempty"`
	// The aggregate amounts calculated per discount for all line items.
	Lines []CreditNoteLineItem `json:"lines,omitempty" database:"foreignKey:id"`
	// Customer-facing text that appears on the credit note PDF.
	Memo string `json:"memo,omitempty"`
	// A unique number that identifies this particular credit note and appears on the PDF of the credit note and its associated invoice.
	Number string `json:"number,omitempty"`
	// Amount that was credited outside of Stripe.
	OutOfBandAmount int `json:"out_of_band_amount,omitempty"`
	// The link to download the PDF of the credit note.
	Pdf string `json:"pdf,omitempty"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason string `json:"reason,omitempty"`
	// Status of this credit note, one of `issued` or `void`. Learn more about [voiding credit notes](https://stripe.com/docs/billing/invoices/credit-notes#voiding).
	Status string `json:"status,omitempty"`
	// The integer amount in %s representing the amount of the credit note, excluding exclusive tax and invoice level discounts.
	Subtotal int `json:"subtotal,omitempty"`
	// The integer amount in %s representing the amount of the credit note, excluding all tax and invoice level discounts.
	SubtotalExcludingTax int `json:"subtotal_excluding_tax,omitempty"`
	// The aggregate amounts calculated per tax rate for all line items.
	TaxAmounts []tax.TaxRate `json:"tax_amounts,omitempty" database:"foreignKey:id"`
	// The integer amount in %s representing the total amount of the credit note, including tax and all discount.
	Total int `json:"total,omitempty"`
	// The integer amount in %s representing the total amount of the credit note, excluding tax, but including discounts.
	TotalExcludingTax int `json:"total_excluding_tax,omitempty"`
	// Type of this credit note, one of `pre_payment` or `post_payment`. A `pre_payment` credit note means it was issued when the invoice was open. A `post_payment` credit note means it was issued when the invoice was paid.
	Type string `json:"type,omitempty"`
	// The time that the credit note was voided.
	VoidedAt time.Time `json:"voided_at,omitempty"`

	CustomerId                   uuid.UUID                   `json:"customer_id" database:"default:null"`
	Customer                     *customer.Customer          `json:"customer,omitempty" database:"foreignKey:customer_id"`
	CustomerBalanceTransactionId uuid.UUID                   `json:"customer_balance_transaction_id,omitempty"`
	CustomerBalanceTransaction   *balance.BalanceTransaction `json:"customer_balance_transaction,omitempty" database:"foreignKey:customer_balance_transaction_id"`
	InvoiceId                    uuid.UUID                   `json:"invoice_id,omitempty"`
	Invoice                      *invoice.Invoice            `json:"invoice,omitempty" database:"foreignKey:invoice_id"`
	RefundId                     uuid.UUID                   `json:"refund_id,omitempty"`
	Refund                       *refund.Refund              `json:"refund,omitempty" database:"foreignKey:refund_id"`
}
