package subscriptionItem

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/tax"
)

// SubscriptionItemBillingThresholds
type SubscriptionItemBillingThresholds struct {
	core.Model

	// Usage threshold that triggers the subscription to create an invoice
	UsageGte int `json:"usage_gte,omitempty"`
}

// SubscriptionItemUpdateParams struct for SubscriptionItemUpdateParams
type SubscriptionItemUpdateParams struct {
	core.Model

	BillingThresholds string `json:"billing_thresholds,omitempty"`
	ClearUsage        bool   `json:"clear_usage,omitempty"`
	Deleted           bool   `json:"deleted,omitempty"`
	Id                string `json:"id,omitempty"`
	Metadata          string `json:"metadata,omitempty"`
	Price             string `json:"price,omitempty"`
	PriceData         string `json:"price_data,omitempty"`
	Quantity          int    `json:"quantity,omitempty"`
	TaxRates          string `json:"tax_rates,omitempty"`
}

// SubscriptionItem Subscription items allow you to create customer subscriptions with more than one plan, making it easy to represent complex billing relationships.
type SubscriptionItem struct {
	core.Model

	BillingThresholds SubscriptionItemBillingThresholds `json:"billing_thresholds,omitempty" database:"foreignKey:id"`
	// Price Price `json:"price"`

	// The [quantity](https://stripe.com/docs/subscriptions/quantities) of the plan to which the customer should be subscribed.
	Quantity int `json:"quantity,omitempty"`
	// The `subscription` this `subscription_item` belongs to.
	Subscription string `json:"subscription"`
	// The tax rates which apply to this `subscription_item`. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.
	TaxRates tax.TaxRate `json:"tax_rates,omitempty" database:"foreignKey:id"`
}
