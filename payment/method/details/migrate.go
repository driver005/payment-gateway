package details

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&PaymentMethodDetails{},
		&ThreeDSecureDetails{},
		&PaymentMethodCardWalletDetails{},
		&PaymentMethodDetailsAchCreditTransfer{},
		&PaymentMethodDetailsAchDebit{},
		&PaymentMethodDetailsAcssDebit{},
		&PaymentMethodDetailsAffirm{},
		&PaymentMethodDetailsAfterpayClearpay{},
		&PaymentMethodDetailsAlipay{},
		&PaymentMethodDetailsAuBecsDebit{},
		&PaymentMethodDetailsBacsDebit{},
		&PaymentMethodDetailsBancontact{},
		&PaymentMethodDetailsBlik{},
		&PaymentMethodDetailsBoleto{},
		&PaymentMethodDetailsCard{},
		&PaymentMethodDetailsCardPresent{},
		&PaymentMethodDetailsCustomerBalance{},
		&PaymentMethodDetailsEps{},
		&PaymentMethodDetailsFpx{},
		&PaymentMethodDetailsGiropay{},
		&PaymentMethodDetailsGrabpay{},
		&PaymentMethodDetailsIdeal{},
		&PaymentMethodDetailsInteracPresent{},
		&PaymentMethodDetailsKlarna{},
		&PaymentMethodDetailsKonbini{},
		&PaymentMethodDetailsLink{},
		&PaymentMethodDetailsMultibanco{},
		&PaymentMethodDetailsOxxo{},
		&PaymentMethodDetailsP24{},
		&PaymentMethodDetailsPaynow{},
		&PaymentMethodDetailsPix{},
		&PaymentMethodDetailsPromptpay{},
		&PaymentMethodDetailsSepaDebit{},
		&PaymentMethodDetailsSofort{},
		&PaymentMethodDetailsStripeAccount{},
		&PaymentMethodDetailsUsBankAccount{},
		&PaymentMethodDetailsWechat{},
		&PaymentMethodDetailsWechatPay{},
	)
	if err != nil {
		panic(err)
	}
}
