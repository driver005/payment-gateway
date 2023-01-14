package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Order - Represents an order
type Order struct {
	core.Model

	// The order's status
	Status string `json:"status" database:"default:null"`

	// The order's fulfillment status
	FulfillmentStatus string `json:"fulfillment_status" database:"default:null"`

	// The order's payment status
	PaymentStatus string `json:"payment_status" database:"default:null"`

	// The order's display ID
	DisplayId int `json:"display_id" database:"default:null"`

	// The ID of the cart associated with the order
	CartId uuid.NullUUID `json:"cart_id" database:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" database:"foreignKey:id;references:cart_id"`

	// The ID of the customer associated with the order
	CustomerId uuid.NullUUID `json:"customer_id"`

	// A customer object. Available if the relation `customer` is expanded.
	Customer *Customer `json:"customer" database:"foreignKey:id;references:customer_id"`

	// The email associated with the order
	Email string `json:"email"`

	// The ID of the billing address associated with the order
	BillingAddressId uuid.NullUUID `json:"billing_address_id" database:"default:null"`

	BillingAddress *Address `json:"billing_address" database:"foreignKey:id;references:billing_address_id"`

	// The ID of the shipping address associated with the order
	ShippingAddressId uuid.NullUUID `json:"shipping_address_id" database:"default:null"`

	ShippingAddress *Address `json:"shipping_address" database:"foreignKey:id;references:shipping_address_id"`

	// The region's ID
	RegionId uuid.NullUUID `json:"region_id"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" database:"foreignKey:id;references:region_id"`

	// The 3 character currency code that is used in the order
	CurrencyCode string `json:"currency_code"`

	Currency *Currency `json:"currency" database:"foreignKey:code;references:currency_code"`

	// The order's tax rate
	TaxRate float64 `json:"tax_rate" database:"default:null"`

	// The discounts used in the order. Available if the relation `discounts` is expanded.
	Discounts []Discount `json:"discounts" database:"foreignKey:id"`

	// The gift cards used in the order. Available if the relation `gift_cards` is expanded.
	GiftCards []GiftCard `json:"gift_cards" database:"foreignKey:id"`

	// The shipping methods used in the order. Available if the relation `shipping_methods` is expanded.
	ShippingMethods []ShippingMethod `json:"shipping_methods" database:"foreignKey:id"`

	// The payments used in the order. Available if the relation `payments` is expanded.
	Payments []Payment `json:"payments" database:"foreignKey:id"`

	// The fulfillments used in the order. Available if the relation `fulfillments` is expanded.
	Fulfillments []Fulfillment `json:"fulfillments" database:"foreignKey:id"`

	// The returns associated with the order. Available if the relation `returns` is expanded.
	Returns []Return `json:"returns" database:"foreignKey:id"`

	// The claims associated with the order. Available if the relation `claims` is expanded.
	Claims []ClaimOrder `json:"claims" database:"foreignKey:id"`

	// The refunds associated with the order. Available if the relation `refunds` is expanded.
	Refunds []Refund `json:"refunds" database:"foreignKey:id"`

	// The swaps associated with the order. Available if the relation `swaps` is expanded.
	Swaps []Swap `json:"swaps" database:"foreignKey:id"`

	// The ID of the draft order this order is associated with.
	DraftOrderId uuid.NullUUID `json:"draft_order_id" database:"default:null"`

	// A draft order object. Available if the relation `draft_order` is expanded.
	DraftOrder *DraftOrder `json:"draft_order" database:"foreignKey:id;references:draft_order_id"`

	// The line items that belong to the order. Available if the relation `items` is expanded.
	Items []LineItem `json:"items" database:"foreignKey:id"`

	// [EXPERIMENTAL] Order edits done on the order. Available if the relation `edits` is expanded.
	Edits []OrderEdit `json:"edits" database:"foreignKey:id"`

	// The gift card transactions used in the order. Available if the relation `gift_card_transactions` is expanded.
	GiftCardTransactions []GiftCardTransaction `json:"gift_card_transactions" database:"foreignKey:id"`

	// The date the order was canceled on.
	CanceledAt time.Time `json:"canceled_at" database:"default:null"`

	// Flag for describing whether or not notifications related to this should be send.
	NoNotification bool `json:"no_notification" database:"default:null"`

	// Randomly generated key used to continue the processing of the order in case of failure.
	IdempotencyKey string `json:"idempotency_key" database:"default:null"`

	// The ID of an external order.
	ExternalId uuid.NullUUID `json:"external_id" database:"default:null"`

	// The ID of the sales channel this order is associated with.
	SalesChannelId uuid.NullUUID `json:"sales_channel_id" database:"default:null"`

	// A sales channel object. Available if the relation `sales_channel` is expanded.
	SalesChannel *SalesChannel `json:"sales_channel" database:"foreignKey:id;references:sales_channel_id"`

	// The total of shipping
	ShippingTotal int `json:"shipping_total" database:"default:null"`

	// The total of discount
	DiscountTotal int `json:"discount_total" database:"default:null"`

	// The total of tax
	TaxTotal int `json:"tax_total" database:"default:null"`

	// The total amount refunded if the order is returned.
	RefundedTotal int `json:"refunded_total" database:"default:null"`

	// The total amount of the order
	Total int `json:"total" database:"default:null"`

	// The subtotal of the order
	Subtotal int `json:"subtotal" database:"default:null"`

	// The total amount paid
	PaidTotal int `json:"paid_total" database:"default:null"`

	// The amount that can be refunded
	RefundableAmount int `json:"refundable_amount" database:"default:null"`

	// The total of gift cards
	GiftCardTotal int `json:"gift_card_total" database:"default:null"`

	// The total of gift cards with taxes
	GiftCardTaxTotal int `json:"gift_card_tax_total" database:"default:null"`
}
