/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostDiscountsDiscountConditionsRequest struct {

	// Operator of the condition
	Operator string `json:"operator"`

	// list of product IDs if the condition is applied on products.
	Products []string `json:"products,omitempty"`

	// list of product type IDs if the condition is applied on product types.
	ProductTypes []string `json:"product_types,omitempty"`

	// list of product collection IDs if the condition is applied on product collections.
	ProductCollections []string `json:"product_collections,omitempty"`

	// list of product tag IDs if the condition is applied on product tags.
	ProductTags []string `json:"product_tags,omitempty"`

	// list of customer group IDs if the condition is applied on customer groups.
	CustomerGroups []string `json:"customer_groups,omitempty"`
}
