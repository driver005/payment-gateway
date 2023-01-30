package options

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&InvoicePaymentMethodOptions{},
		&InvoicePaymentMethodOptionsAcssDebit{},
		&InvoicePaymentMethodOptionsBancontact{},
		&InvoicePaymentMethodOptionsCard{},
		&InvoicePaymentMethodOptionsCustomerBalance{},
		&InvoicePaymentMethodOptionsKonbini{},
		&InvoicePaymentMethodOptionsUsBankAccount{},
	)
	if err != nil {
		panic(err)
	}
}
