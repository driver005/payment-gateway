/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DeleteSalesChannelsChannelProductsBatchRequest struct {

	// The IDs of the products to delete from the Sales Channel.
	ProductIds []DeleteSalesChannelsChannelProductsBatchRequestProductIdsInner `json:"product_ids"`
}