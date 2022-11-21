package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/gofrs/uuid"
)

// ClaimOrder - Claim Orders represent a group of faulty or missing items. Each claim order consists of a subset of items associated with an original order, and can contain additional information about fulfillments and returns.
type ClaimOrder struct {
	core.Model

	Type ClaimStatus `json:"type"`

	// The status of the claim's payment
	PaymentStatus string `json:"payment_status" database:"default:null"`

	FulfillmentStatus string `json:"fulfillment_status" database:"default:null"`

	// The items that have been claimed
	ClaimItems []ClaimItem `json:"claim_items" database:"foreignKey:id"`

	// Refers to the new items to be shipped when the claim order has the type `replace`
	AdditionalItems []LineItem `json:"additional_items" database:"foreignKey:id"`

	// The ID of the order that the claim comes from.
	OrderId uuid.NullUUID `json:"order_id"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" database:"foreignKey:id;references:order_id"`

	// A return object. Holds information about the return if the claim is to be returned. Available if the relation 'return_order' is expanded
	ReturnOrder *Return `json:"return_order" database:"foreignKey:id"`

	// The ID of the address that the new items should be shipped to
	ShippingAddressId uuid.NullUUID `json:"shipping_address_id" database:"default:null"`

	ShippingAddress *Address `json:"shipping_address" database:"foreignKey:id;references:shipping_address_id"`

	// The shipping methods that the claim order will be shipped with.
	ShippingMethods []ShippingMethod `json:"shipping_methods" database:"foreignKey:id"`

	// The fulfillments of the new items to be shipped
	Fulfillments []Fulfillment `json:"fulfillments" database:"foreignKey:id"`

	// The amount that will be refunded in conjunction with the claim
	RefundAmount int32 `json:"refund_amount" database:"default:null"`

	// The date with timezone at which the claim was canceled.
	CanceledAt time.Time `json:"canceled_at" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`

	// Flag for describing whether or not notifications related to this should be send.
	NoNotification bool `json:"no_notification" database:"default:null"`

	// Randomly generated key used to continue the completion of the cart associated with the claim in case of failure.
	IdempotencyKey string `json:"idempotency_key" database:"default:null"`
}
