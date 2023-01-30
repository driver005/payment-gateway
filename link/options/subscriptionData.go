package options

import "github.com/driver005/gateway/core"

// PaymentLinksResourceSubscriptionData
type PaymentLinksResourceSubscriptionData struct {
	core.Model

	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description,omitempty"`
	// Integer representing the number of trial period days before the customer is charged for the first time.
	TrialPeriodDays int `json:"trial_period_days,omitempty"`
}
