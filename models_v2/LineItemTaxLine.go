package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// LineItemTaxLine - Represents an Line Item Tax Line
type LineItemTaxLine struct {
	core.Model

	// The ID of the line item
	ItemId uuid.NullUUID `json:"item_id"`

	Item *LineItem `json:"item" database:"foreignKey:id;references:item_id"`

	// A code to identify the tax type by
	Code string `json:"code" database:"default:null"`

	// A human friendly name for the tax
	Name string `json:"name"`

	// The numeric rate to charge tax by
	Rate float64 `json:"rate"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
