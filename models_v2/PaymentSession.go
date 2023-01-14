package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// PaymentSession - Payment Sessions are created when a Customer initilizes the checkout flow, and can be used to hold the state of a payment flow. Each Payment Session is controlled by a Payment Provider, who is responsible for the communication with external payment services. Authorized Payment Sessions will eventually get promoted to Payments to indicate that they are authorized for capture/refunds/etc.
type PaymentSession struct {
	core.Model

	// The id of the Cart that the Payment Session is created for.
	CartId uuid.NullUUID `json:"cart_id"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" database:"foreignKey:id;references:cart_id"`

	// The id of the Payment Provider that is responsible for the Payment Session
	ProviderId uuid.NullUUID `json:"provider_id"`

	// A flag to indicate if the Payment Session has been selected as the method that will be used to complete the purchase.
	IsSelected bool `json:"is_selected" database:"default:null"`

	// Indicates the status of the Payment Session. Will default to `pending`, and will eventually become `authorized`. Payment Sessions may have the status of `requires_more` to indicate that further actions are to be completed by the Customer.
	Status string `json:"status"`

	// The data required for the Payment Provider to identify, modify and process the Payment Session. Typically this will be an object that holds an id to the external payment session, but can be an empty object if the Payment Provider doesn't hold any state.
	Data JSONB `json:"data" database:"default:null"`

	// Randomly generated key used to continue the completion of a cart in case of failure.
	IdempotencyKey string `json:"idempotency_key" database:"default:null"`
}
