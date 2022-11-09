package models

import "time"

// Money Amounts represents an amount that a given Product Variant can be purcased for. Each Money Amount either has a Currency or Region associated with it to indicate the pricing in a given Currency or, for fully region-based pricing, the given price in a specific Region. If region-based pricing is used the amount will be in the currency defined for the Reigon.
type MoneyAmount struct {
	// The amount in the smallest currecny unit (e.g. cents 100 cents to charge $1) that the Product Variant will cost.
	Amount int `json:"amount"`

	// Currency
	Currency *Currency `json:"currency,omitempty"`

	// The 3 character currency code that the Money Amount is given in.
	CurrencyCode string `json:"currency_code"`

	// The money amount's ID
	Id *string `json:"id,omitempty"`

	// The maximum quantity that the Money Amount applies to. If this value is not set, the Money Amount applies to all quantities.
	MaxQuantity *int `json:"max_quantity,omitempty"`

	// An optional key-value map with additional details
	Metadata *map[string]interface{} `json:"metadata,omitempty"`

	// The minimum quantity that the Money Amount applies to. If this value is not set, the Money Amount applies to all quantities.
	MinQuantity *int `json:"min_quantity,omitempty"`

	// Price Lists represents a set of prices that overrides the default price for one or more product variants.
	PriceList *PriceList `json:"price_list,omitempty"`

	// The ID of the price list associated with the money amount
	PriceListId *string `json:"price_list_id,omitempty"`

	// A region object. Available if the relation `region` is expanded.
	Region *map[string]interface{} `json:"region,omitempty"`

	// The region's ID
	RegionId *string `json:"region_id,omitempty"`

	// The Product Variant contained in the Line Item. Available if the relation `variant` is expanded.
	Variant *map[string]interface{} `json:"variant,omitempty"`

	// The id of the Product Variant contained in the Line Item.
	VariantId *string `json:"variant_id,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
