package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// LineItem - Line Items represent purchasable units that can be added to a Cart for checkout. When Line Items are purchased they will get copied to the resulting order and can eventually be referenced in Fulfillments and Returns. Line Items may also be created when processing Swaps and Claims.
type LineItem struct {
	core.Model

	// The ID of the Cart that the Line Item belongs to.
	CartId uuid.NullUUID `json:"cart_id" database:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" database:"foreignKey:id;references:cart_id"`

	// The ID of the Order that the Line Item belongs to.
	OrderId uuid.NullUUID `json:"order_id" database:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" database:"foreignKey:id;references:order_id"`

	// The id of the Swap that the Line Item belongs to.
	SwapId uuid.NullUUID `json:"swap_id" database:"default:null"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap *Swap `json:"swap" database:"foreignKey:id;references:swap_id"`

	// The id of the Claim that the Line Item belongs to.
	ClaimOrderId uuid.NullUUID `json:"claim_order_id" database:"default:null"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder *ClaimOrder `json:"claim_order" database:"foreignKey:id;references:claim_order_id"`

	// Available if the relation `tax_lines` is expanded.
	TaxLines []LineItemTaxLine `json:"tax_lines" database:"foreignKey:id"`

	// Available if the relation `adjustments` is expanded.
	Adjustments []LineItemAdjustment `json:"adjustments" database:"foreignKey:id"`

	// The title of the Line Item, this should be easily identifiable by the Customer.
	Title string `json:"title"`

	// A more detailed description of the contents of the Line Item.
	Description string `json:"description" database:"default:null"`

	// A URL string to a small image of the contents of the Line Item.
	Thumbnail string `json:"thumbnail" database:"default:null"`

	// Is the item being returned
	IsReturn bool `json:"is_return" database:"default:null"`

	// Flag to indicate if the Line Item is a Gift Card.
	IsGiftcard bool `json:"is_giftcard" database:"default:null"`

	// Flag to indicate if new Line Items with the same variant should be merged or added as an additional Line Item.
	ShouldMerge bool `json:"should_merge" database:"default:null"`

	// Flag to indicate if the Line Item should be included when doing discount calculations.
	AllowDiscounts bool `json:"allow_discounts" database:"default:null"`

	// Flag to indicate if the Line Item has fulfillment associated with it.
	HasShipping bool `json:"has_shipping" database:"default:null"`

	// The price of one unit of the content in the Line Item. This should be in the currency defined by the Cart/Order/Swap/Claim that the Line Item belongs to.
	UnitPrice int `json:"unit_price"`

	// The id of the Product Variant contained in the Line Item.
	VariantId uuid.NullUUID `json:"variant_id" database:"default:null"`

	// A product variant object. The Product Variant contained in the Line Item. Available if the relation `variant` is expanded.
	Variant *ProductVariant `json:"variant" database:"foreignKey:id;references:variant_id"`

	// The quantity of the content in the Line Item.
	Quantity int `json:"quantity"`

	// The quantity of the Line Item that has been fulfilled.
	FulfilledQuantity int `json:"fulfilled_quantity" database:"default:null"`

	// The quantity of the Line Item that has been returned.
	ReturnedQuantity int `json:"returned_quantity" database:"default:null"`

	// The quantity of the Line Item that has been shipped.
	ShippedQuantity int `json:"shipped_quantity" database:"default:null"`

	// The amount that can be refunded from the given Line Item. Takes taxes and discounts into consideration.
	Refundable int `json:"refundable" database:"default:null"`

	// The subtotal of the line item
	Subtotal int `json:"subtotal" database:"default:null"`

	// The total of tax of the line item
	TaxTotal int `json:"tax_total" database:"default:null"`

	// The total amount of the line item
	Total int `json:"total" database:"default:null"`

	// The original total amount of the line item
	OriginalTotal int `json:"original_total" database:"default:null"`

	// The original tax total amount of the line item
	OriginalTaxTotal int `json:"original_tax_total" database:"default:null"`

	// The total of discount of the line item
	DiscountTotal int `json:"discount_total" database:"default:null"`

	// The total of the gift card of the line item
	GiftCardTotal int `json:"gift_card_total" database:"default:null"`

	// [EXPERIMENTAL] Indicates if the line item unit_price include tax
	IncludesTax bool `json:"includes_tax" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
