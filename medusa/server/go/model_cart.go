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

// Cart - Represents a user cart
type Cart struct {

	// The cart's ID
	Id string `json:"id,omitempty"`

	// The email associated with the cart
	Email string `json:"email,omitempty"`

	// The billing address's ID
	BillingAddressId string `json:"billing_address_id,omitempty"`

	BillingAddress Address `json:"billing_address,omitempty"`

	// The shipping address's ID
	ShippingAddressId string `json:"shipping_address_id,omitempty"`

	ShippingAddress Address `json:"shipping_address,omitempty"`

	// Available if the relation `items` is expanded.
	Items []LineItem `json:"items,omitempty"`

	// The region's ID
	RegionId string `json:"region_id,omitempty"`

	// A region object. Available if the relation `region` is expanded.
	Region map[string]interface{} `json:"region,omitempty"`

	// Available if the relation `discounts` is expanded.
	Discounts []map[string]interface{} `json:"discounts,omitempty"`

	// Available if the relation `gift_cards` is expanded.
	GiftCards []map[string]interface{} `json:"gift_cards,omitempty"`

	// The customer's ID
	CustomerId string `json:"customer_id,omitempty"`

	// A customer object. Available if the relation `customer` is expanded.
	Customer map[string]interface{} `json:"customer,omitempty"`

	PaymentSession PaymentSession `json:"payment_session,omitempty"`

	// The payment sessions created on the cart.
	PaymentSessions []PaymentSession `json:"payment_sessions,omitempty"`

	// The payment's ID if available
	PaymentId string `json:"payment_id,omitempty"`

	Payment Payment `json:"payment,omitempty"`

	// The shipping methods added to the cart.
	ShippingMethods []ShippingMethod `json:"shipping_methods,omitempty"`

	// The cart's type.
	Type string `json:"type,omitempty"`

	// The date with timezone at which the cart was completed.
	CompletedAt time.Time `json:"completed_at,omitempty"`

	// The date with timezone at which the payment was authorized.
	PaymentAuthorizedAt time.Time `json:"payment_authorized_at,omitempty"`

	// Randomly generated key used to continue the completion of a cart in case of failure.
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// The context of the cart which can include info like IP or user agent.
	Context map[string]interface{} `json:"context,omitempty"`

	// The sales channel ID the cart is associated with.
	SalesChannelId string `json:"sales_channel_id,omitempty"`

	// A sales channel object. Available if the relation `sales_channel` is expanded.
	SalesChannel map[string]interface{} `json:"sales_channel,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The total of shipping
	ShippingTotal int32 `json:"shipping_total,omitempty"`

	// The total of discount
	DiscountTotal int32 `json:"discount_total,omitempty"`

	// The total of tax
	TaxTotal int32 `json:"tax_total,omitempty"`

	// The total amount refunded if the order associated with this cart is returned.
	RefundedTotal int32 `json:"refunded_total,omitempty"`

	// The total amount of the cart
	Total int32 `json:"total,omitempty"`

	// The subtotal of the cart
	Subtotal int32 `json:"subtotal,omitempty"`

	// The amount that can be refunded
	RefundableAmount int32 `json:"refundable_amount,omitempty"`

	// The total of gift cards
	GiftCardTotal int32 `json:"gift_card_total,omitempty"`

	// The total of gift cards with taxes
	GiftCardTaxTotal int32 `json:"gift_card_tax_total,omitempty"`
}
