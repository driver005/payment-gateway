package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Cart - Represents a user cart
type Cart struct {
	core.Model

	// The email associated with the cart
	Email string `json:"email" database:"default:null"`

	// The billing address's ID
	BillingAddressId uuid.NullUUID `json:"billing_address_id" database:"default:null"`

	BillingAddress *Address `json:"billing_address" database:"foreignKey:id;references:billing_address_id"`

	// The shipping address's ID
	ShippingAddressId uuid.NullUUID `json:"shipping_address_id" database:"default:null"`

	ShippingAddress *Address `json:"shipping_address" database:"foreignKey:id;references:shipping_address_id"`

	// Available if the relation `items` is expanded.
	Items []LineItem `json:"items" database:"foreignKey:id"`

	// The region's ID
	RegionId uuid.NullUUID `json:"region_id" database:"default:null"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" database:"foreignKey:id;references:region_id"`

	// Available if the relation `discounts` is expanded.
	Discounts []Discount `json:"discounts" database:"foreignKey:id"`

	// Available if the relation `gift_cards` is expanded.
	GiftCards []GiftCard `json:"gift_cards" database:"foreignKey:id"`

	// The customer's ID
	CustomerId uuid.NullUUID `json:"customer_id" database:"default:null"`

	// A customer object. Available if the relation `customer` is expanded.
	Customer *Customer `json:"customer" database:"foreignKey:id;references:customer_id"`

	PaymentSession *PaymentSession `json:"payment_session" database:"foreignKey:id"`

	// The payment sessions created on the cart.
	PaymentSessions []PaymentSession `json:"payment_sessions" database:"foreignKey:id"`

	// The payment's ID if available
	PaymentId uuid.NullUUID `json:"payment_id" database:"default:null"`

	Payment *Payment `json:"payment" database:"foreignKey:id;references:payment_id"`

	// The shipping methods added to the cart.
	ShippingMethods []ShippingMethod `json:"shipping_methods" database:"foreignKey:id"`

	// The cart's type.
	Type string `json:"type" database:"default:null"`

	// The date with timezone at which the cart was completed.
	CompletedAt time.Time `json:"completed_at" database:"default:null"`

	// The date with timezone at which the payment was authorized.
	PaymentAuthorizedAt time.Time `json:"payment_authorized_at" database:"default:null"`

	// Randomly generated key used to continue the completion of a cart in case of failure.
	IdempotencyKey string `json:"idempotency_key" database:"default:null"`

	// The context of the cart which can include info like IP or user agent.
	Context JSONB `json:"context" database:"default:null"`

	// The sales channel ID the cart is associated with.
	SalesChannelId uuid.NullUUID `json:"sales_channel_id" database:"default:null"`

	// A sales channel object. Available if the relation `sales_channel` is expanded.
	SalesChannel *SalesChannel `json:"sales_channel" database:"foreignKey:id;references:sales_channel_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`

	// The total of shipping
	ShippingTotal int32 `json:"shipping_total" database:"default:null"`

	// The total of discount
	DiscountTotal int32 `json:"discount_total" database:"default:null"`

	// The total of tax
	TaxTotal int32 `json:"tax_total" database:"default:null"`

	// The total amount refunded if the order associated with this cart is returned.
	RefundedTotal int32 `json:"refunded_total" database:"default:null"`

	// The total amount of the cart
	Total int32 `json:"total" database:"default:null"`

	// The subtotal of the cart
	Subtotal int32 `json:"subtotal" database:"default:null"`

	// The amount that can be refunded
	RefundableAmount int32 `json:"refundable_amount" database:"default:null"`

	// The total of gift cards
	GiftCardTotal int32 `json:"gift_card_total" database:"default:null"`

	// The total of gift cards with taxes
	GiftCardTaxTotal int32 `json:"gift_card_tax_total" database:"default:null"`
}
