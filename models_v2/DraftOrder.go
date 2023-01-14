package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// DraftOrder - Represents a draft order
type DraftOrder struct {
	core.Model

	// The status of the draft order
	Status DraftOrderStatus `json:"status" database:"default:open"`

	// The draft order's display ID
	DisplayId uuid.NullUUID `json:"display_id" database:"default:null"`

	// The ID of the cart associated with the draft order.
	CartId uuid.NullUUID `json:"cart_id" database:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart Cart `json:"cart" database:"default:null"`

	// The ID of the order associated with the draft order.
	OrderId uuid.NullUUID `json:"order_id" database:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" database:"foreignKey:id;references:order_id"`

	// The date the draft order was canceled at.
	CanceledAt time.Time `json:"canceled_at" database:"default:null"`

	// The date the draft order was completed at.
	CompletedAt time.Time `json:"completed_at" database:"default:null"`

	// Whether to send the customer notifications regarding order updates.
	NoNotificationOrder bool `json:"no_notification_order" database:"default:null"`

	// Randomly generated key used to continue the completion of the cart associated with the draft order in case of failure.
	IdempotencyKey string `json:"idempotency_key" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
