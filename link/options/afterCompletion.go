package options

import "github.com/driver005/gateway/core"

// PaymentLinksResourceAfterCompletion
type PaymentLinksResourceAfterCompletion struct {
	core.Model

	HostedConfirmation *PaymentLinksResourceCompletionBehaviorConfirmationPage `json:"hosted_confirmation,omitempty" database:"foreignKey:id"`
	Redirect           *PaymentLinksResourceCompletionBehaviorRedirect         `json:"redirect,omitempty" database:"foreignKey:id"`
	// The specified behavior after the purchase is complete.
	Type string `json:"type,omitempty"`
}
