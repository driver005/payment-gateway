/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostProductsProductVariantsRequestPricesInner struct {

	// The ID of the price.
	Id string `json:"id,omitempty"`

	// The ID of the Region for which the price is used. Only required if currency_code is not provided.
	RegionId string `json:"region_id,omitempty"`

	// The 3 character ISO currency code for which the price will be used. Only required if region_id is not provided.
	CurrencyCode string `json:"currency_code,omitempty"`

	// The amount to charge for the Product Variant.
	Amount int32 `json:"amount"`

	// The minimum quantity for which the price will be used.
	MinQuantity int32 `json:"min_quantity,omitempty"`

	// The maximum quantity for which the price will be used.
	MaxQuantity int32 `json:"max_quantity,omitempty"`
}
