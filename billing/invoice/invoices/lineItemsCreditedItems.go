package invoices

// InvoicesLineItemsCreditedItems
type InvoicesLineItemsCreditedItems struct {
	// Invoice containing the credited invoice line items
	Invoice string `json:"invoice"`
	// Credited invoice line items
	InvoiceLineItems []string `json:"invoice_line_items"`
}
