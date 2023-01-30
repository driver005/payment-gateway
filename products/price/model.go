package price

import (
	"github.com/driver005/gateway/core"
)

type CustomUnitAmount struct {
	core.Model

	Maximum int `json:"maximum,omitempty"`
	Minimum int `json:"minimum,omitempty"`
	Preset  int `json:"preset,omitempty"`
}

type CurrencyOption struct {
	core.Model

	CustomUnitAmount  *CustomUnitAmount `json:"custom_unit_amount,omitempty"`
	TaxBehavior       string            `json:"tax_behavior,omitempty"`
	Tiers             []PriceTier       `json:"tiers,omitempty"`
	UnitAmount        int               `json:"unit_amount,omitempty"`
	UnitAmountDecimal float64           `json:"unit_amount_decimal,omitempty"`
}

type TransformQuantity struct {
	core.Model

	DivideBy int32  `json:"divide_by,omitempty"`
	Round    string `json:"round,omitempty"`
}

type Recurring struct {
	core.Model

	AggregateUsage string `json:"aggregate_usage,omitempty"`
	Interval       string `json:"interval,omitempty"`
	IntervalCount  int32  `json:"interval_count,omitempty"`
	UsageType      string `json:"usage_type,omitempty"`
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
	BillingScheme     string             `json:"billing_scheme,omitempty"`
	Currency          string             `json:"currency,omitempty"`
	CurrencyOptions   *CurrencyOption    `json:"currency_options,omitempty"`
	CustomUnitAmount  *CustomUnitAmount  `json:"custom_unit_amount,omitempty"`
	LookupKey         string             `json:"lookup_key,omitempty"`
	Nickname          string             `json:"nickname,omitempty"`
	Recurring         *Recurring         `json:"recurring,omitempty"`
	TaxBehavior       string             `json:"tax_behavior,omitempty"`
	Tiers             []PriceTier        `json:"tiers,omitempty"`
	TiersMode         string             `json:"tiers_mode,omitempty"`
	TransformQuantity *TransformQuantity `json:"transform_quantity,omitempty"`
	Type              string             `json:"type,omitempty"`
	UnitAmount        int                `json:"unit_amount,omitempty"`
	UnitAmountDecimal float64            `json:"unit_amount_decimal,omitempty"`
}
