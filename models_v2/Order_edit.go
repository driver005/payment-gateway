package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// OrderEdit - Order edit keeps track of order items changes.
type OrderEdit struct {
	core.Model

	// The ID of the order that is edited
	OrderId uuid.NullUUID `json:"order_id"`

	Order *Order `json:"order" database:"foreignKey:id;references:order_id"`

	// Line item changes array.
	Changes []OrderItemChange `json:"changes" database:"foreignKey:id"`

	// An optional note with additional details about the order edit.
	InternalNote string `json:"internal_note" database:"default:null"`

	// The unique identifier of the user or customer who created the order edit.
	CreatedBy string `json:"created_by"`

	// The unique identifier of the user or customer who requested the order edit.
	RequestedBy string `json:"requested_by" database:"default:null"`

	// The date with timezone at which the edit was requested.
	RequestedAt time.Time `json:"requested_at" database:"default:null"`

	// The unique identifier of the user or customer who confirmed the order edit.
	ConfirmedBy string `json:"confirmed_by" database:"default:null"`

	// The date with timezone at which the edit was confirmed.
	ConfirmedAt time.Time `json:"confirmed_at" database:"default:null"`

	// The unique identifier of the user or customer who declined the order edit.
	DeclinedBy string `json:"declined_by" database:"default:null"`

	// The date with timezone at which the edit was declined.
	DeclinedAt time.Time `json:"declined_at" database:"default:null"`

	// An optional note why  the order edit is declined.
	DeclinedReason string `json:"declined_reason" database:"default:null"`

	// The subtotal for line items computed from changes.
	Subtotal int `json:"subtotal" database:"default:null"`

	// The total of discount
	DiscountTotal int `json:"discount_total" database:"default:null"`

	// The total of tax
	TaxTotal int `json:"tax_total" database:"default:null"`

	// The total amount of the edited order.
	Total int `json:"total" database:"default:null"`

	// The difference between the total amount of the order and total amount of edited order.
	DifferenceDue int `json:"difference_due" database:"default:null"`

	// Computed line items from the changes.
	Items []LineItem `json:"items" database:"foreignKey:id"`
}
