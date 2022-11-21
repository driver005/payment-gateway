package models

import (
	"github.com/driver005/gateway/core"
	"github.com/gofrs/uuid"
)

// TaxRate - A Tax Rate can be used to associate a certain rate to charge on products within a given Region
type TaxRate struct {
	core.Model

	// The numeric rate to charge
	Rate float32 `json:"rate" database:"default:null"`

	// A code to identify the tax type by
	Code string `json:"code" database:"default:null"`

	// A human friendly name for the tax
	Name string `json:"name"`

	// The id of the Region that the rate belongs to
	RegionId uuid.NullUUID `json:"region_id"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" database:"foreignKey:id;references:region_id"`

	// The products that belong to this tax rate. Available if the relation `products` is expanded.
	Products []Product `json:"products" database:"foreignKey:id"`

	// The product types that belong to this tax rate. Available if the relation `product_types` is expanded.
	ProductTypes []ProductType `json:"product_types" database:"foreignKey:id"`

	// The shipping options that belong to this tax rate. Available if the relation `shipping_options` is expanded.
	ShippingOptions []ShippingOption `json:"shipping_options" database:"foreignKey:id"`

	// The count of products
	ProductCount int32 `json:"product_count" database:"default:null"`

	// The count of product types
	ProductTypeCount int32 `json:"product_type_count" database:"default:null"`

	// The count of shipping options
	ShippingOptionCount int32 `json:"shipping_option_count" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
