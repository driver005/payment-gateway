package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionWechatPayRedirectToAndroidApp
type PaymentIntentNextActionWechatPayRedirectToAndroidApp struct {
	core.Model

	// app_id is the APP ID registered on WeChat open platform
	AppId string `json:"app_id,omitempty"`
	// nonce_str is a random string
	NonceStr string `json:"nonce_str,omitempty"`
	// package is static value
	Package string `json:"package,omitempty"`
	// an unique merchant ID assigned by WeChat Pay
	PartnerId string `json:"partner_id,omitempty"`
	// an unique trading ID assigned by WeChat Pay
	PrepayId string `json:"prepay_id,omitempty"`
	// A signature
	Sign string `json:"sign,omitempty"`
	// Specifies the current time in epoch format
	Timestamp string `json:"timestamp,omitempty"`
}
