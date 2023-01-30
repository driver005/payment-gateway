package page

import "github.com/driver005/gateway/core"

// PaymentPagesCheckoutSessionInvoiceCreation
type PaymentPagesCheckoutSessionInvoiceCreation struct {
	core.Model

	// Indicates whether invoice creation is enabled for the Checkout Session.
	Enabled     bool                                        `json:"enabled,omitempty"`
	InvoiceData *PaymentPagesCheckoutSessionInvoiceSettings `json:"invoice_data,omitempty" database:"foreignKey:id"`
}
