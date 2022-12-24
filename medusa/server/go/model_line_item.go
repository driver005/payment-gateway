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

// LineItem - Line Items represent purchasable units that can be added to a Cart for checkout. When Line Items are purchased they will get copied to the resulting order and can eventually be referenced in Fulfillments and Returns. Line Items may also be created when processing Swaps and Claims.
type LineItem struct {

	// The cart's ID
	Id string `json:"id,omitempty"`

	// The ID of the Cart that the Line Item belongs to.
	CartId string `json:"cart_id,omitempty"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart map[string]interface{} `json:"cart,omitempty"`

	// The ID of the Order that the Line Item belongs to.
	OrderId string `json:"order_id,omitempty"`

	// An order object. Available if the relation `order` is expanded.
	Order map[string]interface{} `json:"order,omitempty"`

	// The id of the Swap that the Line Item belongs to.
	SwapId string `json:"swap_id,omitempty"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap map[string]interface{} `json:"swap,omitempty"`

	// The id of the Claim that the Line Item belongs to.
	ClaimOrderId string `json:"claim_order_id,omitempty"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder map[string]interface{} `json:"claim_order,omitempty"`

	// Available if the relation `tax_lines` is expanded.
	TaxLines []LineItemTaxLine `json:"tax_lines,omitempty"`

	// Available if the relation `adjustments` is expanded.
	Adjustments []LineItemAdjustment `json:"adjustments,omitempty"`

	// The title of the Line Item, this should be easily identifiable by the Customer.
	Title string `json:"title"`

	// A more detailed description of the contents of the Line Item.
	Description string `json:"description,omitempty"`

	// A URL string to a small image of the contents of the Line Item.
	Thumbnail string `json:"thumbnail,omitempty"`

	// Is the item being returned
	IsReturn bool `json:"is_return,omitempty"`

	// Flag to indicate if the Line Item is a Gift Card.
	IsGiftcard bool `json:"is_giftcard,omitempty"`

	// Flag to indicate if new Line Items with the same variant should be merged or added as an additional Line Item.
	ShouldMerge bool `json:"should_merge,omitempty"`

	// Flag to indicate if the Line Item should be included when doing discount calculations.
	AllowDiscounts bool `json:"allow_discounts,omitempty"`

	// Flag to indicate if the Line Item has fulfillment associated with it.
	HasShipping bool `json:"has_shipping,omitempty"`

	// The price of one unit of the content in the Line Item. This should be in the currency defined by the Cart/Order/Swap/Claim that the Line Item belongs to.
	UnitPrice bool `json:"unit_price"`

	// The id of the Product Variant contained in the Line Item.
	VariantId string `json:"variant_id,omitempty"`

	// A product variant object. The Product Variant contained in the Line Item. Available if the relation `variant` is expanded.
	Variant map[string]interface{} `json:"variant,omitempty"`

	// The quantity of the content in the Line Item.
	Quantity int32 `json:"quantity"`

	// The quantity of the Line Item that has been fulfilled.
	FulfilledQuantity int32 `json:"fulfilled_quantity,omitempty"`

	// The quantity of the Line Item that has been returned.
	ReturnedQuantity int32 `json:"returned_quantity,omitempty"`

	// The quantity of the Line Item that has been shipped.
	ShippedQuantity int32 `json:"shipped_quantity,omitempty"`

	// The amount that can be refunded from the given Line Item. Takes taxes and discounts into consideration.
	Refundable int32 `json:"refundable,omitempty"`

	// The subtotal of the line item
	Subtotal int32 `json:"subtotal,omitempty"`

	// The total of tax of the line item
	TaxTotal int32 `json:"tax_total,omitempty"`

	// The total amount of the line item
	Total int32 `json:"total,omitempty"`

	// The original total amount of the line item
	OriginalTotal int32 `json:"original_total,omitempty"`

	// The original tax total amount of the line item
	OriginalTaxTotal int32 `json:"original_tax_total,omitempty"`

	// The total of discount of the line item
	DiscountTotal int32 `json:"discount_total,omitempty"`

	// The total of the gift card of the line item
	GiftCardTotal int32 `json:"gift_card_total,omitempty"`

	// [EXPERIMENTAL] Indicates if the line item unit_price include tax
	IncludesTax bool `json:"includes_tax,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}