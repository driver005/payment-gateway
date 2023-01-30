package options

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&SubscriptionsResourcePaymentMethodOptions{},
		&SubscriptionsResourcePaymentMethodOptionsAcssDebit{},
		&SubscriptionsResourcePaymentMethodOptionsBancontact{},
		&SubscriptionsResourcePaymentMethodOptionsCard{},
		&SubscriptionsResourcePaymentMethodOptionsCustomerBalance{},
		&SubscriptionsResourcePaymentMethodOptionsKonbini{},
		&SubscriptionsResourcePaymentMethodOptionsUsBankAccount{},
	)
	if err != nil {
		panic(err)
	}
}
