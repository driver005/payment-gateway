package methods

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&PaymentIntentPaymentMethodOptions{},
		&PaymentIntentPaymentMethodOptionsAcssDebit1{},
		&PaymentIntentPaymentMethodOptionsAffirm{},
		&PaymentIntentPaymentMethodOptionsAfterpayClearpay{},
		&PaymentIntentPaymentMethodOptionsAlipay{},
		&PaymentIntentPaymentMethodOptionsAuBecsDebit1{},
		&PaymentIntentPaymentMethodOptionsBacsDebit{},
		&PaymentIntentPaymentMethodOptionsBancontact{},
		&PaymentIntentPaymentMethodOptionsBlik1{},
		&PaymentIntentPaymentMethodOptionsBoleto{},
		&PaymentIntentPaymentMethodOptionsCard1{},
		&PaymentIntentPaymentMethodOptionsCardPresent{},
		&PaymentIntentPaymentMethodOptionsCustomerBalance{},
		&PaymentIntentPaymentMethodOptionsEps1{},
		&PaymentIntentPaymentMethodOptionsFpx{},
		&PaymentIntentPaymentMethodOptionsGiropay{},
		&PaymentIntentPaymentMethodOptionsGrabpay{},
		&PaymentIntentPaymentMethodOptionsIdeal{},
		&PaymentIntentPaymentMethodOptionsInteracPresent{},
		&PaymentIntentPaymentMethodOptionsKlarna{},
		&PaymentIntentPaymentMethodOptionsKonbini{},
		&PaymentIntentPaymentMethodOptionsLink1{},
		&PaymentIntentPaymentMethodOptionsOxxo{},
		&PaymentIntentPaymentMethodOptionsP24{},
		&PaymentIntentPaymentMethodOptionsPaynow{},
		&PaymentIntentPaymentMethodOptionsPix{},
		&PaymentIntentPaymentMethodOptionsPromptpay{},
		&PaymentIntentPaymentMethodOptionsSepaDebit1{},
		&PaymentIntentPaymentMethodOptionsSofort{},
		&PaymentIntentPaymentMethodOptionsUsBankAccount1{},
		&PaymentIntentPaymentMethodOptionsWechatPay{},
	)
	if err != nil {
		panic(err)
	}
}
