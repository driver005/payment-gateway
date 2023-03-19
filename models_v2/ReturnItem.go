package models

import "github.com/google/uuid"

// ReturnItem - Correlates a Line Item with a Return, keeping track of the quantity of the Line Item that will be returned.
type ReturnItem struct {

	// The id of the Return that the Return Item belongs to.
	ReturnId uuid.NullUUID `json:"return_id" database:"primarykey"`

	ReturnOrder *Return `json:"return_order" database:"foreignKey:id;references:return_id"`

	// The id of the Line Item that the Return Item references.
	ItemId uuid.NullUUID `json:"item_id"`

	Item *LineItem `json:"item" database:"foreignKey:id;references:item_id"`

	// The quantity of the Line Item that is included in the Return.
	Quantity int `json:"quantity" database:"default:null"`

	// Whether the Return Item was requested initially or received unexpectedly in the warehouse.
	IsRequested bool `json:"is_requested" database:"default:null"`

	// The quantity that was originally requested to be returned.
	RequestedQuantity int `json:"requested_quantity" database:"default:null"`

	// The quantity that was received in the warehouse.
	RecievedQuantity int `json:"recieved_quantity" database:"default:null"`

	// The ID of the reason for returning the item.
	ReasonId uuid.NullUUID `json:"reason_id,omitempty"`

	Reason *ReturnReason `json:"reason" database:"foreignKey:id;references:reason_id"`

	// An optional note with additional details about the Return.
	Note string `json:"note" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
