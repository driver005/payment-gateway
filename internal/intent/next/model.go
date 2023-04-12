package next

import (
	"github.com/driver005/gateway/core"
)

// PaymentIntentNextAction
type PaymentIntentNextAction struct {
	core.Model

	AlipayHandleRedirect            *PaymentIntentNextActionAlipayHandleRedirect            `json:"alipay_handle_redirect,omitempty" database:"foreignKey:id"`
	BoletoDisplayDetails            *PaymentIntentNextActionBoleto                          `json:"boleto_display_details,omitempty" database:"foreignKey:id"`
	CardAwaitNotification           *PaymentIntentNextActionCardAwaitNotification           `json:"card_await_notification,omitempty" database:"foreignKey:id"`
	DisplayBankTransferInstructions *PaymentIntentNextActionDisplayBankTransferInstructions `json:"display_bank_transfer_instructions,omitempty" database:"foreignKey:id"`
	KonbiniDisplayDetails           *PaymentIntentNextActionKonbini                         `json:"konbini_display_details,omitempty" database:"foreignKey:id"`
	OxxoDisplayDetails              *PaymentIntentNextActionDisplayOxxoDetails              `json:"oxxo_display_details,omitempty" database:"foreignKey:id"`
	PaynowDisplayQrCode             *PaymentIntentNextActionPaynowDisplayQrCode             `json:"paynow_display_qr_code,omitempty" database:"foreignKey:id"`
	PixDisplayQrCode                *PaymentIntentNextActionPixDisplayQrCode                `json:"pix_display_qr_code,omitempty" database:"foreignKey:id"`
	PromptpayDisplayQrCode          *PaymentIntentNextActionPromptpayDisplayQrCode          `json:"promptpay_display_qr_code,omitempty" database:"foreignKey:id"`
	RedirectToUrl                   *PaymentIntentNextActionRedirectToUrl                   `json:"redirect_to_url,omitempty" database:"foreignKey:id"`
	VerifyWithMicrodeposits         *PaymentIntentNextActionVerifyWithMicrodeposits         `json:"verify_with_microdeposits,omitempty" database:"foreignKey:id"`
	WechatPayDisplayQrCode          *PaymentIntentNextActionWechatPayDisplayQrCode          `json:"wechat_pay_display_qr_code,omitempty" database:"foreignKey:id"`
	WechatPayRedirectToAndroidApp   *PaymentIntentNextActionWechatPayRedirectToAndroidApp   `json:"wechat_pay_redirect_to_android_app,omitempty" database:"foreignKey:id"`
	WechatPayRedirectToIosApp       *PaymentIntentNextActionWechatPayRedirectToIosApp       `json:"wechat_pay_redirect_to_ios_app,omitempty" database:"foreignKey:id"`
	UseStripeSdk                    core.JSONB                                              `json:"use_stripe_sdk,omitempty"`
	Type                            string                                                  `json:"type,omitempty"`
}
