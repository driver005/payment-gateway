package next

// PaymentIntentNextActionWechatPayRedirectToAndroidApp
type PaymentIntentNextActionWechatPayRedirectToAndroidApp struct {
	// app_id is the APP ID registered on WeChat open platform
	AppId string `json:"app_id"`
	// nonce_str is a random string
	NonceStr string `json:"nonce_str"`
	// package is static value
	Package string `json:"package"`
	// an unique merchant ID assigned by WeChat Pay
	PartnerId string `json:"partner_id"`
	// an unique trading ID assigned by WeChat Pay
	PrepayId string `json:"prepay_id"`
	// A signature
	Sign string `json:"sign"`
	// Specifies the current time in epoch format
	Timestamp string `json:"timestamp"`
}
