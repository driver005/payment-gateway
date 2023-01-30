package options

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&PaymentLinksResourceAfterCompletion{},
		&PaymentLinksResourceCompletionBehaviorConfirmationPage{},
		&PaymentLinksResourceCompletionBehaviorRedirect{},
		&PaymentLinksResourceConsentCollection{},
		&PaymentLinksResourceCustomText{},
		&PaymentLinksResourcePaymentIntentData{},
		&PaymentLinksResourceShippingAddressCollection{},
		&PaymentLinksResourceShippingOption{},
		&PaymentLinksResourceSubscriptionData{},
	)
	if err != nil {
		panic(err)
	}
}
