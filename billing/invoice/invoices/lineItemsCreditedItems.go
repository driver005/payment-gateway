package invoices

import "github.com/driver005/gateway/core"

// InvoicesLineItemsCreditedItems
type InvoicesLineItemsCreditedItems struct {
	core.Model
	// Invoice containing the credited invoice line items
	Invoice string `json:"invoice,omitempty"`
	// Credited invoice line items
	InvoiceLineItems []string `json:"invoice_line_items,omitempty" database:"type:text[]"`
}
