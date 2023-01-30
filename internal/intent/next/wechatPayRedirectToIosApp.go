package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionWechatPayRedirectToIosApp
type PaymentIntentNextActionWechatPayRedirectToIosApp struct {
	core.Model

	// An universal link that redirect to WeChat Pay app
	NativeUrl string `json:"native_url,omitempty"`
}
