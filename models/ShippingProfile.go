package models

import "time"

// Shipping Profiles have a set of defined Shipping Options that can be used to fulfill a given set of Products.
type ShippingProfile struct {
	// The shipping profile's ID
	Id *string `json:"id,omitempty"`

	// An optional key-value map with additional details
	Metadata *map[string]interface{} `json:"metadata,omitempty"`

	// The name given to the Shipping profile - this may be displayed to the Customer.
	Name string `json:"name"`

	// The Products that the Shipping Profile defines Shipping Options for. Available if the relation `products` is expanded.
	Products *[]map[string]interface{} `json:"products,omitempty"`

	// The Shipping Options that can be used to fulfill the Products in the Shipping Profile. Available if the relation `shipping_options` is expanded.
	ShippingOptions *[]ShippingOption `json:"shipping_options,omitempty"`

	// The type of the Shipping Profile, may be `default`, `gift_card` or `custom`.
	Type ShippingProfileType `json:"type"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
