package next

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&PaymentIntentNextAction{},
		&PaymentIntentNextActionAlipayHandleRedirect{},
		&PaymentIntentNextActionBoleto{},
		&PaymentIntentNextActionCardAwaitNotification{},
		&PaymentIntentNextActionDisplayBankTransferInstructions{},
		&PaymentIntentNextActionKonbini{},
		&PaymentIntentNextActionDisplayOxxoDetails{},
		&PaymentIntentNextActionPaynowDisplayQrCode{},
		&PaymentIntentNextActionPixDisplayQrCode{},
		&PaymentIntentNextActionPromptpayDisplayQrCode{},
		&PaymentIntentNextActionRedirectToUrl{},
		&PaymentIntentNextActionVerifyWithMicrodeposits{},
		&PaymentIntentNextActionWechatPayDisplayQrCode{},
		&PaymentIntentNextActionWechatPayRedirectToAndroidApp{},
		&PaymentIntentNextActionWechatPayRedirectToIosApp{},
	)
	if err != nil {
		panic(err)
	}
}
