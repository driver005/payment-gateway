package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type PriceListPricesInner struct {
	core.Model

	// The 3 character currency code that the Money Amount is given in.
	CurrencyCode string `json:"currency_code"`

	Currency *Currency `json:"currency" database:"foreignKey:code;references:currency_code"`

	// The amount in the smallest currecny unit (e.g. cents 100 cents to charge $1) that the Product Variant will cost.
	Amount int32 `json:"amount"`

	// The minimum quantity that the Money Amount applies to. If this value is not set, the Money Amount applies to all quantities.
	MinQuantity int32 `json:"min_quantity" database:"default:null"`

	// The maximum quantity that the Money Amount applies to. If this value is not set, the Money Amount applies to all quantities.
	MaxQuantity int32 `json:"max_quantity" database:"default:null"`

	// The ID of the price list associated with the money amount
	PriceListId uuid.NullUUID `json:"price_list_id" database:"default:null"`

	PriceList *PriceList `json:"price_list" database:"foreignKey:id;references:price_list_id"`

	// The id of the Product Variant contained in the Line Item.
	VariantId uuid.NullUUID `json:"variant_id" database:"default:null"`

	// The Product Variant contained in the Line Item. Available if the relation `variant` is expanded.
	Variant *ProductVariant `json:"variant" database:"foreignKey:id;references:variant_id"`

	// The region's ID
	RegionId uuid.NullUUID `json:"region_id" database:"default:null"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" database:"foreignKey:id;references:region_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`

	// The name of the customer group
	Name string `json:"name"`

	// The customers that belong to the customer group. Available if the relation `customers` is expanded.
	Customers []Customer `json:"customers" database:"foreignKey:id"`

	// The price lists that are associated with the customer group. Available if the relation `price_lists` is expanded.
	PriceLists []PriceList `json:"price_lists" database:"foreignKey:id"`
}
