package next

import (
	"github.com/driver005/gateway/core"
)

// PaymentIntentNextAction
type PaymentIntentNextAction struct {
	core.Model

	AlipayHandleRedirect            *PaymentIntentNextActionAlipayHandleRedirect            `json:"alipay_handle_redirect,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BoletoDisplayDetails            *PaymentIntentNextActionBoleto                          `json:"boleto_display_details,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CardAwaitNotification           *PaymentIntentNextActionCardAwaitNotification           `json:"card_await_notification,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	DisplayBankTransferInstructions *PaymentIntentNextActionDisplayBankTransferInstructions `json:"display_bank_transfer_instructions,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	KonbiniDisplayDetails           *PaymentIntentNextActionKonbini                         `json:"konbini_display_details,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	OxxoDisplayDetails              *PaymentIntentNextActionDisplayOxxoDetails              `json:"oxxo_display_details,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	PaynowDisplayQrCode             *PaymentIntentNextActionPaynowDisplayQrCode             `json:"paynow_display_qr_code,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	PixDisplayQrCode                *PaymentIntentNextActionPixDisplayQrCode                `json:"pix_display_qr_code,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	PromptpayDisplayQrCode          *PaymentIntentNextActionPromptpayDisplayQrCode          `json:"promptpay_display_qr_code,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	RedirectToUrl                   *PaymentIntentNextActionRedirectToUrl                   `json:"redirect_to_url,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	VerifyWithMicrodeposits         *PaymentIntentNextActionVerifyWithMicrodeposits         `json:"verify_with_microdeposits,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	WechatPayDisplayQrCode          *PaymentIntentNextActionWechatPayDisplayQrCode          `json:"wechat_pay_display_qr_code,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	WechatPayRedirectToAndroidApp   *PaymentIntentNextActionWechatPayRedirectToAndroidApp   `json:"wechat_pay_redirect_to_android_app,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	WechatPayRedirectToIosApp       *PaymentIntentNextActionWechatPayRedirectToIosApp       `json:"wechat_pay_redirect_to_ios_app,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UseStripeSdk                    core.JSONB                                              `json:"use_stripe_sdk,omitempty"`
	Type                            string                                                  `json:"type,omitempty"`
}
