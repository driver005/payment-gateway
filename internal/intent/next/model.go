package next

// PaymentIntentNextAction
type PaymentIntentNextAction struct {
	AlipayHandleRedirect            *PaymentIntentNextActionAlipayHandleRedirect            `json:"alipay_handle_redirect,omitempty"`
	BoletoDisplayDetails            *PaymentIntentNextActionBoleto                          `json:"boleto_display_details,omitempty"`
	CardAwaitNotification           *PaymentIntentNextActionCardAwaitNotification           `json:"card_await_notification,omitempty"`
	DisplayBankTransferInstructions *PaymentIntentNextActionDisplayBankTransferInstructions `json:"display_bank_transfer_instructions,omitempty"`
	KonbiniDisplayDetails           *PaymentIntentNextActionKonbini                         `json:"konbini_display_details,omitempty"`
	OxxoDisplayDetails              *PaymentIntentNextActionDisplayOxxoDetails              `json:"oxxo_display_details,omitempty"`
	PaynowDisplayQrCode             *PaymentIntentNextActionPaynowDisplayQrCode             `json:"paynow_display_qr_code,omitempty"`
	PixDisplayQrCode                *PaymentIntentNextActionPixDisplayQrCode                `json:"pix_display_qr_code,omitempty"`
	PromptpayDisplayQrCode          *PaymentIntentNextActionPromptpayDisplayQrCode          `json:"promptpay_display_qr_code,omitempty"`
	RedirectToUrl                   *PaymentIntentNextActionRedirectToUrl                   `json:"redirect_to_url,omitempty"`
	// Type of the next action to perform, one of `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.
	Type string `json:"type"`
	// When confirming a PaymentIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.
	UseStripeSdk                  map[string]interface{}                                `json:"use_stripe_sdk,omitempty"`
	VerifyWithMicrodeposits       *PaymentIntentNextActionVerifyWithMicrodeposits       `json:"verify_with_microdeposits,omitempty"`
	WechatPayDisplayQrCode        *PaymentIntentNextActionWechatPayDisplayQrCode        `json:"wechat_pay_display_qr_code,omitempty"`
	WechatPayRedirectToAndroidApp *PaymentIntentNextActionWechatPayRedirectToAndroidApp `json:"wechat_pay_redirect_to_android_app,omitempty"`
	WechatPayRedirectToIosApp     *PaymentIntentNextActionWechatPayRedirectToIosApp     `json:"wechat_pay_redirect_to_ios_app,omitempty"`
}
