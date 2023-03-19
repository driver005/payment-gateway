package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Return - Return orders hold information about Line Items that a Customer wishes to send back, along with how the items will be returned. Returns can be used as part of a Swap.
type Return struct {
	core.Model

	// Status of the Return.
	Status string `json:"status" database:"default:null"`

	// The Return Items that will be shipped back to the warehouse. Available if the relation `items` is expanded.
	Items []ReturnItem `json:"items" database:"foreignKey:return_id"`

	// The ID of the Swap that the Return is a part of.
	SwapId uuid.NullUUID `json:"swap_id,omitempty"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap *Swap `json:"swap" database:"foreignKey:id;references:swap_id"`

	// The ID of the Order that the Return is made from.
	OrderId uuid.NullUUID `json:"order_id,omitempty"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" database:"foreignKey:id;references:order_id"`

	// The ID of the Claim that the Return is a part of.
	ClaimOrderId uuid.NullUUID `json:"claim_order_id,omitempty"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder *ClaimOrder `json:"claim_order" database:"foreignKey:id;references:claim_order_id"`

	// The Shipping Method that will be used to send the Return back. Can be null if the Customer facilitates the return shipment themselves. Available if the relation `shipping_method` is expanded.
	ShippingMethod []ShippingMethod `json:"shipping_method" database:"foreignKey:id"`

	// Data about the return shipment as provided by the Fulfilment Provider that handles the return shipment.
	ShippingData JSONB `json:"shipping_data" database:"default:null"`

	// The amount that should be refunded as a result of the return.
	RefundAmount int `json:"refund_amount"`

	// When set to true, no notification will be sent related to this return.
	NoNotification bool `json:"no_notification" database:"default:null"`

	// Randomly generated key used to continue the completion of the return in case of failure.
	IdempotencyKey string `json:"idempotency_key" database:"default:null"`

	// The date with timezone at which the return was received.
	ReceivedAt time.Time `json:"received_at" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
