package invoices

import "github.com/driver005/gateway/core"

// InvoicesLineItemsProrationDetails
type InvoicesLineItemsProrationDetails struct {
	core.Model
	CreditedItems *InvoicesLineItemsCreditedItems `json:"credited_items,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
