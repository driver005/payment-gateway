package models

import "time"

// Price Lists represents a set of prices that overrides the default price for one or more product variants.
type PriceList struct {
	// The Customer Groups that the Price List applies to. Available if the relation `customer_groups` is expanded.
	CustomerGroups *[]map[string]interface{} `json:"customer_groups,omitempty"`

	// The price list's description
	Description string `json:"description"`

	// The date with timezone that the Price List stops being valid.
	EndsAt *time.Time `json:"ends_at,omitempty"`

	// The price list's ID
	Id *string `json:"id,omitempty"`

	// [EXPERIMENTAL] Does the price list prices include tax
	IncludesTax *bool `json:"includes_tax,omitempty"`

	// The price list's name
	Name string `json:"name"`

	// The Money Amounts that are associated with the Price List. Available if the relation `prices` is expanded.
	Prices *[]interface{} `json:"prices,omitempty"`

	// The date with timezone that the Price List starts being valid.
	StartsAt *time.Time `json:"starts_at,omitempty"`

	// The status of the Price List
	Status *PriceListStatus `json:"status,omitempty"`

	// The type of Price List. This can be one of either `sale` or `override`.
	Type *PriceListType `json:"type,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
