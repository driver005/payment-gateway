package options

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&CheckoutAcssDebitMandateOptions{},
		&CheckoutAcssDebitPaymentMethodOptions{},
		&CheckoutAffirmPaymentMethodOptions{},
		&CheckoutAfterpayClearpayPaymentMethodOptions{},
		&CheckoutAlipayPaymentMethodOptions{},
		&CheckoutAuBecsDebitPaymentMethodOptions{},
		&CheckoutBacsDebitPaymentMethodOptions{},
		&CheckoutBancontactPaymentMethodOptions{},
		&CheckoutBoletoPaymentMethodOptions{},
		&CheckoutCardInstallmentsOptions{},
		&CheckoutCardPaymentMethodOptions{},
		&CheckoutCustomerBalanceBankTransferPaymentMethodOptions{},
		&CheckoutCustomerBalancePaymentMethodOptions{},
		&CheckoutEpsPaymentMethodOptions{},
		&CheckoutFpxPaymentMethodOptions{},
		&CheckoutGiropayPaymentMethodOptions{},
		&CheckoutGrabPayPaymentMethodOptions{},
		&CheckoutIdealPaymentMethodOptions{},
		&CheckoutKlarnaPaymentMethodOptions{},
		&CheckoutKonbiniPaymentMethodOptions{},
		&CheckoutOxxoPaymentMethodOptions{},
		&CheckoutP24PaymentMethodOptions{},
		&CheckoutPaynowPaymentMethodOptions{},
		&CheckoutPixPaymentMethodOptions{},
		&CheckoutSepaDebitPaymentMethodOptions{},
		&CheckoutSessionPaymentMethodOptions{},
		&CheckoutSofortPaymentMethodOptions{},
		&CheckoutUsBankAccountPaymentMethodOptions{},
	)
	if err != nil {
		panic(err)
	}
}
