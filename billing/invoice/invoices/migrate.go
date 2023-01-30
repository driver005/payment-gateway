package invoices

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&InvoicesLineItemsCreditedItems{},
		&InvoicesLineItemsProrationDetails{},
		&InvoicesPaymentSettings{},
		&InvoicesResourceInvoiceTaxId{},
		&InvoicesStatusTransitions{},
	)
	if err != nil {
		panic(err)
	}
}
