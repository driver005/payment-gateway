/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetCurrencies200Response struct {

	// The number of Currency.
	Count int32 `json:"count,omitempty"`

	// The offset of the Currency query.
	Offset int32 `json:"offset,omitempty"`

	// The limit of the currency query.
	Limit int32 `json:"limit,omitempty"`

	Currencies []Currency `json:"currencies,omitempty"`
}