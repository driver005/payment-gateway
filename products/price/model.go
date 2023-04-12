package price

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/products/product"
	"github.com/google/uuid"
)

type TransformQuantity struct {
	core.Model

	DivideBy int   `json:"divide_by,omitempty"`
	Round    Round `json:"round,omitempty"`
}

type Recurring struct {
	core.Model

	AggregateUsage AggregateUsage `json:"aggregate_usage,omitempty"`
	Interval       Interval       `json:"interval,omitempty"`
	IntervalCount  int            `json:"interval_count,omitempty"`
	UsageType      UsageType      `json:"usage_type,omitempty"`
}

type PriceTier struct {
	core.Model

	FlatAmount        int     `json:"flat_amount,omitempty"`
	FlatAmountDecimal float64 `json:"flat_amount_decimal,omitempty"`
	UnitAmount        int     `json:"unit_amount,omitempty"`
	UnitAmountDecimal float64 `json:"unit_amount_decimal,omitempty"`
	UpTo              int     `json:"up_to,omitempty"`
}

type Price struct {
	core.Model

	Active            bool               `json:"active,omitempty"`
	BillingScheme     BillingScheme      `json:"billing_scheme,omitempty"`
	Currency          string             `json:"currency,omitempty"`
	LookupKey         string             `json:"lookup_key,omitempty"`
	Nickname          string             `json:"nickname,omitempty"`
	Recurring         *Recurring         `json:"recurring,omitempty" database:"foreignKey:id"`
	TaxBehavior       TaxBehavior        `json:"tax_behavior,omitempty"`
	Tiers             []PriceTier        `json:"tiers,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	TiersMode         TiersMode          `json:"tiers_mode,omitempty"`
	TransformQuantity *TransformQuantity `json:"transform_quantity,omitempty" database:"foreignKey:id"`
	UnitAmount        int                `json:"unit_amount,omitempty"`
	UnitAmountDecimal float64            `json:"unit_amount_decimal,omitempty"`
	Type              Type               `json:"type,omitempty"`

	ProductId *uuid.UUID       `json:"product_id,omitempty" swaggerignore:"true"`
	Product   *product.Product `json:"product,omitempty" database:"foreignKey:product_id"`
}
