package options

import "github.com/driver005/gateway/core"

// PaymentLinksResourceCompletionBehaviorConfirmationPage
type PaymentLinksResourceCompletionBehaviorConfirmationPage struct {
	core.Model

	// The custom message that is displayed to the customer after the purchase is complete.
	CustomMessage string `json:"custom_message,omitempty"`
}
