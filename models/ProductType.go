package models

import "github.com/driver005/gateway/core"

// Product Type can be added to Products for filtering and reporting purposes.
type ProductType struct {
	core.Model

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`

	// The value that the Product Type represents.
	Value string `json:"value"`
}
