package page

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&PaymentPagesCheckoutSessionAfterExpiration{},
		&PaymentPagesCheckoutSessionConsent{},
		&PaymentPagesCheckoutSessionConsentCollection{},
		&PaymentPagesCheckoutSessionCustomText{},
		&PaymentPagesCheckoutSessionCustomerDetails{},
		&PaymentPagesCheckoutSessionInvoiceCreation{},
		&PaymentPagesCheckoutSessionInvoiceSettings{},
		&PaymentPagesCheckoutSessionShippingAddressCollection{},
		&PaymentPagesCheckoutSessionShippingCost{},
		&PaymentPagesCheckoutSessionShippingOption{},
		&PaymentPagesCheckoutSessionTaxId{},
		&PaymentPagesCheckoutSessionTotalDetails{},
	)
	if err != nil {
		panic(err)
	}
}
