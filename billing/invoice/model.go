package invoice

import (
	"github.com/driver005/gateway/billing/invoice/invoices"
	"github.com/driver005/gateway/billing/invoice/settings"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/products/item"
	"github.com/driver005/gateway/utils/address"
	"github.com/driver005/gateway/utils/errors"
	"github.com/driver005/gateway/utils/tax"
	"github.com/google/uuid"
)

// InvoiceThresholdReason
type InvoiceThresholdReason struct {
	core.Model

	// The total invoice amount threshold boundary if it triggered the threshold invoice.
	AmountGte int `json:"amount_gte,omitempty"`
	// The IDs of the line items that triggered the threshold invoice.
	LineItemIds []string `json:"line_item_ids,omitempty" database:"type:text[]"`
	// The quantity threshold boundary that applied to the given line item.
	UsageGte int `json:"usage_gte,omitempty"`
}

type Invoice struct {
	core.Model

	AccountCountry               string                                   `json:"account_country,omitempty"`
	AccountName                  string                                   `json:"account_name,omitempty"`
	AmountDue                    int                                      `json:"amount_due,omitempty"`
	AmountPaid                   int                                      `json:"amount_paid,omitempty"`
	AmountRemaining              int                                      `json:"amount_remaining,omitempty"`
	AttemptCount                 int                                      `json:"attempt_count,omitempty"`
	Attempted                    bool                                     `json:"attempted,omitempty"`
	AutoAdvance                  bool                                     `json:"auto_advance,omitempty"`
	AutomaticTax                 tax.AutomaticTax                         `json:"automatic_tax,omitempty" database:"foreignKey:id"`
	BillingReason                string                                   `json:"billing_reason,omitempty"`
	CollectionMethod             string                                   `json:"collection_method,omitempty"`
	Currency                     string                                   `json:"currency,omitempty"`
	CustomFields                 []settings.InvoiceSettingCustomField     `json:"custom_fields,omitempty" database:"foreignKey:id"`
	CustomerAddress              *address.Address                         `json:"customer_address,omitempty" database:"foreignKey:id"`
	CustomerEmail                string                                   `json:"customer_email,omitempty"`
	CustomerName                 string                                   `json:"customer_name,omitempty"`
	CustomerPhone                string                                   `json:"customer_phone,omitempty"`
	CustomerShipping             *address.Shipping                        `json:"customer_shipping,omitempty" database:"foreignKey:id"`
	CustomerTaxExempt            string                                   `json:"customer_tax_exempt,omitempty"`
	CustomerTaxIds               []tax.TaxId                              `json:"customer_tax_ids,omitempty" database:"foreignKey:id"`
	DefaultTaxRates              *tax.TaxRate                             `json:"default_tax_rates,omitempty" database:"foreignKey:id"`
	Description                  string                                   `json:"description,omitempty"`
	DueDate                      int                                      `json:"due_date,omitempty"`
	EndingBalance                int                                      `json:"ending_balance,omitempty"`
	Footer                       string                                   `json:"footer,omitempty"`
	FromInvoice                  *Invoice                                 `json:"from_invoice,omitempty" database:"foreignKey:id"`
	HostedInvoiceUrl             string                                   `json:"hosted_invoice_url,omitempty"`
	InvoicePdf                   string                                   `json:"invoice_pdf,omitempty"`
	LastFinalizationError        *errors.ApiErrors                        `json:"last_finalization_error,omitempty" database:"foreignKey:id"`
	Lines                        []item.LineItem                          `json:"lines,omitempty"`
	NextPaymentAttempt           int                                      `json:"next_payment_attempt,omitempty"`
	Number                       string                                   `json:"number,omitempty"`
	Paid                         bool                                     `json:"paid,omitempty"`
	PaidOutOfBand                bool                                     `json:"paid_out_of_band,omitempty"`
	PaymentSettings              *invoices.InvoicesPaymentSettings        `json:"payment_settings,omitempty" database:"foreignKey:id"`
	PeriodEnd                    int                                      `json:"period_end,omitempty"`
	PeriodStart                  int                                      `json:"period_start,omitempty"`
	PostPaymentCreditNotesAmount int                                      `json:"post_payment_credit_notes_amount,omitempty" database:"foreignKey:id"`
	PrePaymentCreditNotesAmount  int                                      `json:"pre_payment_credit_notes_amount,omitempty" database:"foreignKey:id"`
	ReceiptNumber                string                                   `json:"receipt_number,omitempty"`
	RenderingOptions             *settings.InvoiceSettingRenderingOptions `json:"rendering_options,omitempty"`
	StartingBalance              int                                      `json:"starting_balance,omitempty"`
	StatementDescriptor          string                                   `json:"statement_descriptor,omitempty"`
	Status                       string                                   `json:"status,omitempty"`
	StatusTransitions            *invoices.InvoicesStatusTransitions      `json:"status_transitions,omitempty" database:"foreignKey:id"`
	SubscriptionProrationDate    int                                      `json:"subscription_proration_date,omitempty"`
	Subtotal                     int                                      `json:"subtotal,omitempty"`
	SubtotalExcludingTax         int                                      `json:"subtotal_excluding_tax,omitempty"`
	Tax                          int                                      `json:"tax,omitempty"`
	ThresholdReason              *InvoiceThresholdReason                  `json:"threshold_reason,omitempty" database:"foreignKey:id"`
	Total                        int                                      `json:"total,omitempty"`
	TotalExcludingTax            int                                      `json:"total_excluding_tax,omitempty"`
	TotalTaxAmounts              []tax.TaxRate                            `json:"total_tax_amounts,omitempty" database:"foreignKey:id"`
	WebhooksDeliveredAt          int                                      `json:"webhooks_delivered_at,omitempty"`

	AccountTaxIds *tax.TaxId          `json:"account_tax_ids,omitempty" database:"foreignKey:id"`
	Discounts     []discount.Discount `json:"discounts,omitempty" database:"foreignKey:id"`

	CustomerId             uuid.UUID                  `json:"customer_id" database:"default:null"`
	Customer               *customer.Customer         `json:"customer,omitempty" database:"foreignKey:customer_id"`
	ChargeId               uuid.UUID                  `json:"charge_id,omitempty"`
	Charge                 *charge.Charge             `json:"charge,omitempty" database:"foreignKey:charge_id"`
	DefaultPaymentMethodId uuid.UUID                  `json:"default_payment_method_id" database:"default:null"`
	DefaultPaymentMethod   *method.PaymentMethod      `json:"default_payment_method,omitempty" database:"foreignKey:payment_method_id"`
	PaymentIntentId        uuid.UUID                  `json:"payment_intent_id,omitempty"`
	PaymentIntent          *intent.PaymentIntent      `json:"payment_intent,omitempty" database:"foreignKey:payment_intent_id"`
	LatestRevisionId       uuid.UUID                  `json:"latest_revision_id,omitempty"`
	LatestRevision         *Invoice                   `json:"latest_revision,omitempty" database:"foreignKey:latest_revision_id"`
	SubscriptionId         uuid.UUID                  `json:"subscription_id,omitempty"`
	Subscription           *subscription.Subscription `json:"subscription,omitempty" database:"foreignKey:subscription_id"`
}
