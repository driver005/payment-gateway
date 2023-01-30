package usageRecord

import (
	"github.com/driver005/gateway/core"
)

type Period struct {
	core.Model

	// The end date of this usage period. All usage up to and including this point in time is included.
	End int `json:"end,omitempty"`
	// The start date of this usage period. All usage after this point in time is included.
	Start int `json:"start,omitempty"`
}

// UsageRecordSummary
type UsageRecordSummary struct {
	core.Model

	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// The invoice in which this usage period has been billed for.
	Invoice string `json:"invoice,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	Period Period `json:"period,omitempty" database:"foreignKey:id"`
	// The ID of the subscription item this summary is describing.
	SubscriptionItem string `json:"subscription_item,omitempty"`
	// The total usage within this usage period.
	TotalUsage int `json:"total_usage,omitempty"`
}

// UsageRecord Usage records allow you to report customer usage and metrics to Stripe for metered billing of subscription prices.  Related guide: [Metered Billing](https://stripe.com/docs/billing/subscriptions/metered-billing).
type UsageRecord struct {
	core.Model

	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// The usage quantity for the specified date.
	Quantity int `json:"quantity,omitempty"`
	// The ID of the subscription item this usage record contains data for.
	SubscriptionItem string `json:"subscription_item,omitempty"`
	// The timestamp when this usage occurred.
	Timestamp int `json:"timestamp,omitempty"`
}
