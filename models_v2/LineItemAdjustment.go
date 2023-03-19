package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// LineItemAdjustment - Represents an Line Item Adjustment
type LineItemAdjustment struct {
	core.Model

	// The ID of the line item
	ItemId uuid.NullUUID `json:"item_id"`

	Item *LineItem `json:"item" database:"foreignKey:id;references:item_id"`

	// The line item's adjustment description
	Description string `json:"description"`

	// The ID of the discount associated with the adjustment
	DiscountId uuid.NullUUID `json:"discount_id,omitempty"`

	Discount *Discount `json:"discount" database:"foreignKey:id;references:discount_id"`

	// The adjustment amount
	Amount float64 `json:"amount"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
