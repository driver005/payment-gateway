/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// LineItemAdjustment - Represents an Line Item Adjustment
type LineItemAdjustment struct {

	// The invite's ID
	Id string `json:"id,omitempty"`

	// The ID of the line item
	ItemId string `json:"item_id"`

	Item LineItem `json:"item,omitempty"`

	// The line item's adjustment description
	Description string `json:"description"`

	// The ID of the discount associated with the adjustment
	DiscountId string `json:"discount_id,omitempty"`

	Discount Discount `json:"discount,omitempty"`

	// The adjustment amount
	Amount float32 `json:"amount"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
