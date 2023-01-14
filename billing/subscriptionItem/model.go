package subscriptionItem

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

// SubscriptionItemBillingThresholds
type SubscriptionItemBillingThresholds struct {
	// Usage threshold that triggers the subscription to create an invoice
	UsageGte int `json:"usage_gte,omitempty"`
}

// SubscriptionItemUpdateParams struct for SubscriptionItemUpdateParams
type SubscriptionItemUpdateParams struct {
	BillingThresholds *string `json:"billing_thresholds,omitempty"`
	ClearUsage *bool `json:"clear_usage,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	Id *string `json:"id,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	Price *string `json:"price,omitempty"`
	PriceData *string `json:"price_data,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	TaxRates *string `json:"tax_rates,omitempty"`
}

// SubscriptionItem Subscription items allow you to create customer subscriptions with more than one plan, making it easy to represent complex billing relationships.
type SubscriptionItem struct {
	BillingThresholds SubscriptionItemBillingThresholds `json:"billing_thresholds,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int32 `json:"created"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata core.JSONB `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	
	// Price Price `json:"price"`
	
	// The [quantity](https://stripe.com/docs/subscriptions/quantities) of the plan to which the customer should be subscribed.
	Quantity *int32 `json:"quantity,omitempty"`
	// The `subscription` this `subscription_item` belongs to.
	Subscription string `json:"subscription"`
	// The tax rates which apply to this `subscription_item`. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.
	TaxRates []models.TaxRate `json:"tax_rates,omitempty"`
}
