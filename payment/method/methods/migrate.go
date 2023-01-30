package methods

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&PaymentMethodAcssDebit{},
		&PaymentMethodAffirm{},
		&PaymentMethodAfterpayClearpay{},
		&PaymentMethodAlipay{},
		&PaymentMethodAuBecsDebit{},
		&PaymentMethodBacsDebit{},
		&PaymentMethodBancontact{},
		&PaymentMethodBlik{},
		&PaymentMethodBoleto{},
		&PaymentMethodBtcPay{},
		&PaymentMethodCard{},
		&PaymentMethodCardPresent{},
		&PaymentMethodCustomerBalance{},
		&PaymentMethodEps{},
		&PaymentMethodFpx{},
		&PaymentMethodGiropay{},
		&PaymentMethodGrabpay{},
		&PaymentMethodIdeal{},
		&PaymentMethodInteracPresent{},
		&PaymentMethodKlarna{},
		&PaymentMethodKonbini{},
		&PaymentMethodLink{},
		&PaymentMethodOxxo{},
		&PaymentMethodP24{},
		&PaymentMethodPaynow{},
		&PaymentMethodPix{},
		&PaymentMethodPromptpay{},
		&PaymentMethodSepaDebit{},
		&PaymentMethodSofort{},
		&PaymentMethodUsBankAccount{},
		&PaymentMethodWechatPay{},
	)
	if err != nil {
		panic(err)
	}
}
