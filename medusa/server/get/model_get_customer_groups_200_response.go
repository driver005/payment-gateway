/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetCustomerGroups200Response struct {

	CustomerGroups []CustomerGroup `json:"customer_groups,omitempty"`

	// The total number of items available
	Count int32 `json:"count,omitempty"`

	// The number of items skipped before these items
	Offset int32 `json:"offset,omitempty"`

	// The number of items per page
	Limit int32 `json:"limit,omitempty"`
}