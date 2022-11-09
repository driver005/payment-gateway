package models

import "time"

// Product Options define properties that may vary between different variants of a Product. Common Product Options are "Size" and "Color", but Medusa doesn't limit what Product Options that can be defined.
type ProductOption struct {

	// The product option's ID
	Id *string `json:"id,omitempty"`

	// An optional key-value map with additional details
	Metadata *map[string]interface{} `json:"metadata,omitempty"`

	// A product object. Available if the relation `product` is expanded.
	Product *map[string]interface{} `json:"product,omitempty"`

	// The ID of the Product that the Product Option is defined for.
	ProductId string `json:"product_id"`

	// The title that the Product Option is defined by (e.g. "Size").
	Title string `json:"title"`

	// The Product Option Values that are defined for the Product Option. Available if the relation `values` is expanded.
	Values *[]ProductOptionValue `json:"values,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
