package next

// PaymentIntentNextActionAlipayHandleRedirect
type PaymentIntentNextActionAlipayHandleRedirect struct {
	// The native data to be used with Alipay SDK you must redirect your customer to in order to authenticate the payment in an Android App.
	NativeData string `json:"native_data,omitempty"`
	// The native URL you must redirect your customer to in order to authenticate the payment in an iOS App.
	NativeUrl string `json:"native_url,omitempty"`
	// If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.
	ReturnUrl string `json:"return_url,omitempty"`
	// The URL you must redirect your customer to in order to authenticate the payment.
	Url string `json:"url,omitempty"`
}
