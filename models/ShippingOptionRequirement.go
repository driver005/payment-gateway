package models

import "time"

// A requirement that a Cart must satisfy for the Shipping Option to be available to the Cart.
type ShippingOptionRequirement struct {
	// The amount to compare the Cart subtotal to.
	Amount int `json:"amount"`

	// The shipping option requirement's ID
	Id *string `json:"id,omitempty"`

	// Shipping Options represent a way in which an Order or Return can be shipped. Shipping Options have an associated Fulfillment Provider that will be used when the fulfillment of an Order is initiated. Shipping Options themselves cannot be added to Carts, but serve as a template for Shipping Methods. This distinction makes it possible to customize individual Shipping Methods with additional information.
	ShippingOption *ShippingOption `json:"shipping_option,omitempty"`

	// The id of the Shipping Option that the hipping option requirement belongs to
	ShippingOptionId string `json:"shipping_option_id"`

	// The type of the requirement, this defines how the value will be compared to the Cart's total. `min_subtotal` requirements define the minimum subtotal that is needed for the Shipping Option to be available, while the `max_subtotal` defines the maximum subtotal that the Cart can have for the Shipping Option to be available.
	Type ShippingOptionRequirementType `json:"type"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
