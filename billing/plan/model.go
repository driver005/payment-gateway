package plan

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/products/product"
	"github.com/google/uuid"
)

// TransformUsage
type TransformUsage struct {
	core.Model

	// Divide usage by this number.
	DivideBy int `json:"divide_by,omitempty"`
	// After division, either round the result `up` or `down`.
	Round string `json:"round,omitempty"`
}

// PlanTier
type PlanTier struct {
	core.Model

	// Price for the entire tier.
	FlatAmount int `json:"flat_amount,omitempty"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,omitempty"`
	// Per unit price for units relevant to the tier.
	UnitAmount int `json:"unit_amount,omitempty"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,omitempty"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int `json:"up_to,omitempty"`
}

// Plan You can now model subscriptions more flexibly using the [Prices API](https://stripe.com/docs/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.  Plans define the base price, currency, and billing cycle for recurring purchases of products. [Products](https://stripe.com/docs/api#products) help you track inventory or provisioning, and plans help you track pricing. Different physical goods or levels of service should be represented by products, and pricing options should be represented by plans. This approach lets you change prices without having to change your provisioning scheme.  For example, you might have a single \"gold\" product that has plans for $10/month, $100/year, €9/month, and €90/year.  Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription) and more about [products and prices](https://stripe.com/docs/products-prices/overview).
type Plan struct {
	core.Model

	// Whether the plan can be used for new purchases.
	Active bool `json:"active,omitempty"`
	// Specifies a usage aggregation strategy for plans of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.
	AggregateUsage string `json:"aggregate_usage,omitempty"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	Amount int `json:"amount,omitempty"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	AmountDecimal float64 `json:"amount_decimal,omitempty"`
	// Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `amount`) will be charged per unit in `quantity` (for plans with `usage_type=licensed`), or per unit of total usage (for plans with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.
	BillingScheme string `json:"billing_scheme,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
	Interval string `json:"interval,omitempty"`
	// The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.
	IntervalCount int `json:"interval_count,omitempty"`
	// A brief description of the plan, hidden from customers.
	Nickname string `json:"nickname,omitempty"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []PlanTier `json:"tiers,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	// Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.
	TiersMode      string          `json:"tiers_mode,omitempty"`
	TransformUsage *TransformUsage `json:"transform_usage,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Default number of trial days when subscribing a customer to this plan using [`trial_from_plan=true`](https://stripe.com/docs/api#create_subscription-trial_from_plan).
	TrialPeriodDays int `json:"trial_period_days,omitempty"`
	// Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.
	UsageType string `json:"usage_type,omitempty"`

	ProductId *uuid.UUID       `json:"product_id,omitempty" swaggerignore:"true"`
	Product   *product.Product `json:"product,omitempty" database:"foreignKey:product_id" swaggertype:"primitive,string" format:"uuid"`
}
