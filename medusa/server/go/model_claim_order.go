/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ClaimOrder - Claim Orders represent a group of faulty or missing items. Each claim order consists of a subset of items associated with an original order, and can contain additional information about fulfillments and returns.
type ClaimOrder struct {

	// The claim's ID
	Id string `json:"id,omitempty"`

	Type string `json:"type"`

	// The status of the claim's payment
	PaymentStatus string `json:"payment_status,omitempty"`

	FulfillmentStatus string `json:"fulfillment_status,omitempty"`

	// The items that have been claimed
	ClaimItems []ClaimItem `json:"claim_items,omitempty"`

	// Refers to the new items to be shipped when the claim order has the type `replace`
	AdditionalItems []LineItem `json:"additional_items,omitempty"`

	// The ID of the order that the claim comes from.
	OrderId string `json:"order_id"`

	// An order object. Available if the relation `order` is expanded.
	Order map[string]interface{} `json:"order,omitempty"`

	// A return object. Holds information about the return if the claim is to be returned. Available if the relation 'return_order' is expanded
	ReturnOrder map[string]interface{} `json:"return_order,omitempty"`

	// The ID of the address that the new items should be shipped to
	ShippingAddressId string `json:"shipping_address_id,omitempty"`

	ShippingAddress Address `json:"shipping_address,omitempty"`

	// The shipping methods that the claim order will be shipped with.
	ShippingMethods []ShippingMethod `json:"shipping_methods,omitempty"`

	// The fulfillments of the new items to be shipped
	Fulfillments []Fulfillment `json:"fulfillments,omitempty"`

	// The amount that will be refunded in conjunction with the claim
	RefundAmount int32 `json:"refund_amount,omitempty"`

	// The date with timezone at which the claim was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Flag for describing whether or not notifications related to this should be send.
	NoNotification bool `json:"no_notification,omitempty"`

	// Randomly generated key used to continue the completion of the cart associated with the claim in case of failure.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
