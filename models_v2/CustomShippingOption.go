package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// CustomShippingOption - Custom Shipping Options are 'overriden' Shipping Options. Store managers can attach a Custom Shipping Option to a cart in order to set a custom price for a particular Shipping Option
type CustomShippingOption struct {
	core.Model

	// The custom price set that will override the shipping option's original price
	Price int `json:"price"`

	// The ID of the Shipping Option that the custom shipping option overrides
	ShippingOptionId uuid.NullUUID `json:"shipping_option_id"`

	// A shipping option object. Available if the relation `shipping_option` is expanded.
	ShippingOption *ShippingOption `json:"shipping_option" database:"foreignKey:id;references:shipping_option_id"`

	// The ID of the Cart that the custom shipping option is attached to
	CartId uuid.NullUUID `json:"cart_id" database:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" database:"foreignKey:id;references:cart_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`

	// [EXPERIMENTAL] Indicates if the custom shipping option price include tax
	IncludesTax bool `json:"includes_tax" database:"default:null"`
}
