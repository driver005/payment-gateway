package models

import "time"

// A value given to a Product Variant's option set. Product Variant have a Product Option Value for each of the Product Options defined on the Product.
type ProductOptionValue struct {

	// The product option value's ID
	Id *string `json:"id,omitempty"`

	// An optional key-value map with additional details
	Metadata *map[string]interface{} `json:"metadata,omitempty"`

	// Product Options define properties that may vary between different variants of a Product. Common Product Options are "Size" and "Color", but Medusa doesn't limit what Product Options that can be defined.
	Option *ProductOption `json:"option,omitempty"`

	// The ID of the Product Option that the Product Option Value is defined for.
	OptionId string `json:"option_id"`

	// The value that the Product Variant has defined for the specific Product Option (e.g. if the Product Option is "Size" this value could be "Small", "Medium" or "Large").
	Value string `json:"value"`

	// Product Variants represent a Product with a specific set of Product Option configurations. The maximum number of Product Variants that a Product can have is given by the number of available Product Option combinations.
	Variant *ProductVariant `json:"variant,omitempty"`

	// The ID of the Product Variant that the Product Option Value is defined for.
	VariantId string `json:"variant_id"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
