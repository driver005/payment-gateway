package subscriptionItem

import (
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/utils/tax"
	"github.com/google/uuid"
)

// SubscriptionItem Subscription items allow you to create customer subscriptions with more than one plan, making it easy to represent complex billing relationships.
type SubscriptionItem struct {
	core.Model

	// The [quantity](https://stripe.com/docs/subscriptions/quantities) of the plan to which the customer should be subscribed.
	Quantity int `json:"quantity,omitempty"`
	// The tax rates which apply to this `subscription_item`. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.
	TaxRates *tax.TaxRate `json:"tax_rates,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`

	PriceId        *uuid.UUID                 `json:"price_id,omitempty" swaggerignore:"true"`
	Price          *price.Price               `json:"price,omitempty" database:"foreignKey:price_id" swaggertype:"primitive,string" format:"uuid"`
	SubscriptionId *uuid.UUID                 `json:"subscription_id,omitempty" swaggerignore:"true"`
	Subscription   *subscription.Subscription `json:"subscription,omitempty" database:"foreignKey:subscription_id" swaggertype:"primitive,string" format:"uuid"`
}
