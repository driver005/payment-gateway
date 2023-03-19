package invoices

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// InvoicesLineItemsCreditedItems
type InvoicesLineItemsCreditedItems struct {
	core.Model
	// Invoice containing the credited invoice line items
	Invoice string `json:"invoice,omitempty"`
	// Credited invoice line items
	InvoiceLineItems pq.StringArray `json:"invoice_line_items,omitempty" database:"type:varchar(64)[]"`
}
