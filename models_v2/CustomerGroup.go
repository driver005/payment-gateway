package models

import "github.com/driver005/gateway/core"

// CustomerGroup - Represents a customer group
type CustomerGroup struct {
	core.Model

	// The name of the customer group
	Name string `json:"name"`

	// The customers that belong to the customer group. Available if the relation `customers` is expanded.
	Customers []Customer `json:"customers" database:"many2many:customer_group"`

	// The price lists that are associated with the customer group. Available if the relation `price_lists` is expanded.
	PriceLists []PriceList `json:"price_lists" database:"foreignKey:id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
