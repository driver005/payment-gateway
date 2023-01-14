package usageRecord

type Period struct {
	// The end date of this usage period. All usage up to and including this point in time is included.
	End int `json:"end,omitempty"`
	// The start date of this usage period. All usage after this point in time is included.
	Start int `json:"start,omitempty"`
}

// UsageRecordSummary
type UsageRecordSummary struct {
	// Unique identifier for the object.
	Id string `json:"id"`
	// The invoice in which this usage period has been billed for.
	Invoice string `json:"invoice,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	Period Period `json:"period"`
	// The ID of the subscription item this summary is describing.
	SubscriptionItem string `json:"subscription_item"`
	// The total usage within this usage period.
	TotalUsage int32 `json:"total_usage"`
}

// UsageRecord Usage records allow you to report customer usage and metrics to Stripe for metered billing of subscription prices.  Related guide: [Metered Billing](https://stripe.com/docs/billing/subscriptions/metered-billing).
type UsageRecord struct {
	// Unique identifier for the object.
	Id string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The usage quantity for the specified date.
	Quantity int32 `json:"quantity"`
	// The ID of the subscription item this usage record contains data for.
	SubscriptionItem string `json:"subscription_item"`
	// The timestamp when this usage occurred.
	Timestamp int32 `json:"timestamp"`
}
