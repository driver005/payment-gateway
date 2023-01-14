package invoice

import (
	"github.com/driver005/gateway/billing/customerTax"
	"github.com/driver005/gateway/billing/invoice/invoices"
	"github.com/driver005/gateway/billing/invoice/settings"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/utils/address"
)

// InvoiceThresholdReason
type InvoiceThresholdReason struct {
	// The total invoice amount threshold boundary if it triggered the threshold invoice.
	AmountGte int `json:"amount_gte,omitempty"`
	// Indicates which line items triggered a threshold invoice.
	ItemReasons []InvoiceItemThresholdReason `json:"item_reasons"`
}

// InvoiceTaxAmount
type InvoiceTaxAmount struct {
	// The amount, in %s, of the tax.
	Amount int `json:"amount"`
	// Whether this tax amount is inclusive or exclusive.
	Inclusive bool           `json:"inclusive"`
	TaxRate   models.TaxRate `json:"tax_rate"`
}

// InvoiceMandateOptionsCard
type InvoiceMandateOptionsCard struct {
	// Amount to be charged for future payments.
	Amount int `json:"amount,omitempty"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType string `json:"amount_type,omitempty"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description string `json:"description,omitempty"`
}

// InvoiceLineItemPeriod
type InvoiceLineItemPeriod struct {
	// The end of the period, which must be greater than or equal to the start.
	End int `json:"end"`
	// The start of the period.
	Start int `json:"start"`
}

// InvoiceItemThresholdReason
type InvoiceItemThresholdReason struct {
	// The IDs of the line items that triggered the threshold invoice.
	LineItemIds []string `json:"line_item_ids"`
	// The quantity threshold boundary that applied to the given line item.
	UsageGte int `json:"usage_gte"`
}

// Period2 struct for Period2
type Period2 struct {
	End   int32 `json:"end"`
	Start int32 `json:"start"`
}

// OneTimePriceData1 struct for OneTimePriceData1
type OneTimePriceData struct {
	Currency          string   `json:"currency"`
	Product           string   `json:"product"`
	TaxBehavior       *string  `json:"tax_behavior,omitempty"`
	UnitAmount        *int32   `json:"unit_amount,omitempty"`
	UnitAmountDecimal *float64 `json:"unit_amount_decimal,omitempty"`
}

// InvoiceItemPreviewParams struct for InvoiceItemPreviewParams
type InvoiceItemPreviewParams struct {
	Amount            *int              `json:"amount,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	Description       string            `json:"description,omitempty"`
	Discountable      *bool             `json:"discountable,omitempty"`
	Discounts         *string           `json:"discounts,omitempty"`
	Invoiceitem       string            `json:"invoiceitem,omitempty"`
	Metadata          *core.JSONB       `json:"metadata,omitempty"`
	Period            *Period2          `json:"period,omitempty"`
	Price             string            `json:"price,omitempty"`
	PriceData         *OneTimePriceData `json:"price_data,omitempty"`
	Quantity          *int              `json:"quantity,omitempty"`
	TaxBehavior       string            `json:"tax_behavior,omitempty"`
	TaxCode           *string           `json:"tax_code,omitempty"`
	TaxRates          string            `json:"tax_rates,omitempty"`
	UnitAmount        *int              `json:"unit_amount,omitempty"`
	UnitAmountDecimal *float64          `json:"unit_amount_decimal,omitempty"`
}

// InvoiceDefaultSource ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription's default source, if any, or to the customer's default source.
type InvoiceDefaultSource struct {
	BankAccount *models.BankAccount
	Card        *models.Card
	Source      *models.Source
}

type Invoice struct {
	AccountCountry       string                                  `json:"account_country,omitempty"`
	AccountName          string                                  `json:"account_name,omitempty"`
	AccountTaxIds        []customerTax.TaxId                     `json:"account_tax_ids,omitempty"`
	AmountDue            int                                     `json:"amount_due"`
	AmountPaid           int                                     `json:"amount_paid"`
	AmountRemaining      int                                     `json:"amount_remaining"`
	AttemptCount         int                                     `json:"attempt_count"`
	Attempted            bool                                    `json:"attempted"`
	AutoAdvance          *bool                                   `json:"auto_advance,omitempty"`
	AutomaticTax         models.AutomaticTax                     `json:"automatic_tax"`
	BillingReason        string                                  `json:"billing_reason,omitempty"`
	Charge               *charge.Charge                          `json:"charge,omitempty"`
	CollectionMethod     string                                  `json:"collection_method"`
	Created              int                                     `json:"created"`
	Currency             string                                  `json:"currency"`
	CustomFields         []settings.InvoiceSettingCustomField    `json:"custom_fields,omitempty"`
	Customer             customer.Customer                       `json:"customer,omitempty"`
	CustomerAddress      address.Address                         `json:"customer_address,omitempty"`
	CustomerEmail        string                                  `json:"customer_email,omitempty"`
	CustomerName         string                                  `json:"customer_name,omitempty"`
	CustomerPhone        string                                  `json:"customer_phone,omitempty"`
	CustomerShipping     address.Shipping                        `json:"customer_shipping,omitempty"`
	CustomerTaxExempt    string                                  `json:"customer_tax_exempt,omitempty"`
	CustomerTaxIds       []invoices.InvoicesResourceInvoiceTaxId `json:"customer_tax_ids,omitempty"`
	DefaultPaymentMethod method.PaymentMethod                    `json:"default_payment_method,omitempty"`
	DefaultSource        InvoiceDefaultSource                    `json:"default_source,omitempty"`
	DefaultTaxRates      []models.TaxRate                        `json:"default_tax_rates"`
	Description          string                                  `json:"description,omitempty"`
	// Discount                     Discount                      `json:"discount,omitempty"`
	// Discounts                    []Discount              `json:"discounts,omitempty"`
	DueDate               int              `json:"due_date,omitempty"`
	EndingBalance         int              `json:"ending_balance,omitempty"`
	Footer                string           `json:"footer,omitempty"`
	FromInvoice           *Invoice         `json:"from_invoice,omitempty"`
	HostedInvoiceUrl      string           `json:"hosted_invoice_url,omitempty"`
	Id                    string           `json:"id,omitempty"`
	InvoicePdf            string           `json:"invoice_pdf,omitempty"`
	LastFinalizationError models.ApiErrors `json:"last_finalization_error,omitempty"`
	LatestRevision        *Invoice         `json:"latest_revision,omitempty"`
	// Lines                        []LineItem                    `json:"lines"`
	Livemode           bool              `json:"livemode"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	NextPaymentAttempt int               `json:"next_payment_attempt,omitempty"`
	Number             string            `json:"number,omitempty"`
	Object             string            `json:"object"`
	Paid               bool              `json:"paid"`
	PaidOutOfBand      bool              `json:"paid_out_of_band"`
	// PaymentIntent                intent.PaymentIntent             `json:"payment_intent,omitempty"`
	PaymentSettings              invoices.InvoicesPaymentSettings `json:"payment_settings"`
	PeriodEnd                    int                              `json:"period_end"`
	PeriodStart                  int                              `json:"period_start"`
	PostPaymentCreditNotesAmount int                              `json:"post_payment_credit_notes_amount"`
	PrePaymentCreditNotesAmount  int                              `json:"pre_payment_credit_notes_amount"`
	// Quote                        *Quote                         `json:"quote,omitempty"`
	ReceiptNumber string `json:"receipt_number,omitempty"`
	// RenderingOptions             *InvoiceSettingRenderingOptions              `json:"rendering_options,omitempty"`
	StartingBalance     int                                `json:"starting_balance"`
	StatementDescriptor string                             `json:"statement_descriptor,omitempty"`
	Status              string                             `json:"status,omitempty"`
	StatusTransitions   invoices.InvoicesStatusTransitions `json:"status_transitions"`
	// Subscription                 Subscriptions                  `json:"subscription,omitempty"`
	SubscriptionProrationDate int                     `json:"subscription_proration_date,omitempty"`
	Subtotal                  int                     `json:"subtotal"`
	SubtotalExcludingTax      int                     `json:"subtotal_excluding_tax,omitempty"`
	Tax                       int                     `json:"tax,omitempty"`
	ThresholdReason           *InvoiceThresholdReason `json:"threshold_reason,omitempty"`
	Total                     int                     `json:"total"`
	// TotalDiscountAmounts         []Discount    `json:"total_discount_amounts,omitempty"`
	TotalExcludingTax   int                `json:"total_excluding_tax,omitempty"`
	TotalTaxAmounts     []InvoiceTaxAmount `json:"total_tax_amounts"`
	WebhooksDeliveredAt int                `json:"webhooks_delivered_at,omitempty"`
}
