package rule

import "github.com/driver005/gateway/core"

// Rule
type Rule struct {
	core.Model

	// The action taken on the payment.
	Action string `json:"action,omitempty"`
	// The predicate to evaluate the payment against.
	Predicate string `json:"predicate,omitempty"`
}
