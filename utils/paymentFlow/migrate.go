package paymentFlow

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&PaymentFlowsAmountDetails{},
		&PaymentFlowsAutomaticPaymentMethodsPaymentIntent{},
		&PaymentFlowsInstallmentOptions{},
		&PaymentFlowsPrivatePaymentMethodsAlipayDetails{},
		&PaymentFlowsPrivatePaymentMethodsKlarnaDob{},
	)
	if err != nil {
		panic(err)
	}
}
