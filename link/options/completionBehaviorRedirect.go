package options

import "github.com/driver005/gateway/core"

// PaymentLinksResourceCompletionBehaviorRedirect
type PaymentLinksResourceCompletionBehaviorRedirect struct {
	core.Model

	// The URL the customer will be redirected to after the purchase is complete.
	Url string `json:"url,omitempty"`
}
