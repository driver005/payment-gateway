package options

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&PaymentMethodOptionsAffirm{},
		&PaymentMethodOptionsAfterpayClearpay{},
		&PaymentMethodOptionsAlipay{},
		&PaymentMethodOptionsBacsDebit{},
		&PaymentMethodOptionsBancontact{},
		&PaymentMethodOptionsBoleto{},
		&PaymentMethodOptionsCardInstallments{},
		&PaymentMethodOptionsCardMandateOptions{},
		&PaymentMethodOptionsCardPresent{},
		&PaymentMethodOptionsCustomerBalance{},
		&PaymentMethodOptionsCustomerBalanceBankTransfer{},
		&PaymentMethodOptionsFpx{},
		&PaymentMethodOptionsGiropay{},
		&PaymentMethodOptionsGrabpay{},
		&PaymentMethodOptionsIdeal{},
		&PaymentMethodOptionsKlarna{},
		&PaymentMethodOptionsKonbini{},
		&PaymentMethodOptionsOxxo{},
		&PaymentMethodOptionsP24{},
		&PaymentMethodOptionsPaynow{},
		&PaymentMethodOptionsPix{},
		&PaymentMethodOptionsPromptpay{},
		&PaymentMethodOptionsSofort{},
		&PaymentMethodOptionsWechatPay{},
	)
	if err != nil {
		panic(err)
	}
}
