package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

// Price Lists represents a set of prices that overrides the default price for one or more product variants.
type PriceList struct {
	core.Model

	// The Customer Groups that the Price List applies to. Available if the relation `customer_groups` is expanded.
	CustomerGroups *CustomerGroup `json:"customer_groups" database:"foreignKey:id"`

	// The price list's description
	Description string `json:"description"`

	// The date with timezone that the Price List stops being valid.
	EndsAt time.Time `json:"ends_at" database:"default:null"`

	// [EXPERIMENTAL] Does the price list prices include tax
	IncludesTax bool `json:"includes_tax" database:"default:null"`

	// The price list's name
	Name string `json:"name"`

	// The Money Amounts that are associated with the Price List. Available if the relation `prices` is expanded.
	Prices []MoneyAmount `json:"prices" database:"foreignKey:id"`

	// The date with timezone that the Price List starts being valid.
	StartsAt time.Time `json:"starts_at" database:"default:null"`

	// The status of the Price List
	Status PriceListStatus `json:"status" database:"default:null"`

	// The type of Price List. This can be one of either `sale` or `override`.
	Type PriceListType `json:"type" database:"default:null"`
}
