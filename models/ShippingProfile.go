package models

import (
	"github.com/driver005/gateway/core"
)

// Shipping Profiles have a set of defined Shipping Options that can be used to fulfill a given set of Products.
type ShippingProfile struct {
	core.Model

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`

	// The name given to the Shipping profile - this may be displayed to the Customer.
	Name string `json:"name"`

	// The Products that the Shipping Profile defines Shipping Options for. Available if the relation `products` is expanded.
	Products []Product `json:"products" database:"foreignKey:id"`

	// The Shipping Options that can be used to fulfill the Products in the Shipping Profile. Available if the relation `shipping_options` is expanded.
	ShippingOptions []ShippingOption `json:"shipping_options" database:"foreignKey:id"`

	// The type of the Shipping Profile, may be `default`, `gift_card` or `custom`.
	Type ShippingProfileType `json:"type"`
}

// func (m *ShippingProfile) BeforeCreate(tx *database.DB) (err error) {
// 	m.Id, err = uuid.NewV4()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
