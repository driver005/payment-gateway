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

// Refund - Refund represent an amount of money transfered back to the Customer for a given reason. Refunds may occur in relation to Returns, Swaps and Claims, but can also be initiated by a store operator.
type Refund struct {

	// The refund's ID
	Id string `json:"id,omitempty"`

	// The id of the Order that the Refund is related to.
	OrderId string `json:"order_id"`

	// The amount that has be refunded to the Customer.
	Amount int32 `json:"amount"`

	// An optional note explaining why the amount was refunded.
	Note string `json:"note,omitempty"`

	// The reason given for the Refund, will automatically be set when processed as part of a Swap, Claim or Return.
	Reason string `json:"reason,omitempty"`

	// Randomly generated key used to continue the completion of the refund in case of failure.
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}