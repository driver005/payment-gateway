package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// ReturnReason - A Reason for why a given product is returned. A Return Reason can be used on Return Items in order to indicate why a Line Item was returned.
type ReturnReason struct {
	core.Model

	// A description of the Reason.
	Description string `json:"description" database:"default:null"`

	// A text that can be displayed to the Customer as a reason.
	Label string `json:"label"`

	// The value to identify the reason by.
	Value string `json:"value"`

	// The ID of the parent reason.
	ParentReturnReasonId uuid.NullUUID `json:"parent_return_reason_id,omitempty"`

	ParentReturnReason *ReturnReason `json:"parent_return_reason" database:"foreignKey:id;references:parent_return_reason_id"`

	ReturnReasonChildren *ReturnReason `json:"return_reason_children" database:"foreignKey:id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
