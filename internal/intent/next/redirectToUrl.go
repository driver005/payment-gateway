package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionRedirectToUrl
type PaymentIntentNextActionRedirectToUrl struct {
	core.Model

	// If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.
	ReturnUrl string `json:"return_url,omitempty"`
	// The URL you must redirect your customer to in order to authenticate the payment.
	Url string `json:"url,omitempty"`
}
