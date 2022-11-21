/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostDraftOrdersDraftOrderRequest struct {

	// The ID of the Region to create the Draft Order in.
	RegionId string `json:"region_id,omitempty"`

	// The 2 character ISO code for the Country.
	CountryCode string `json:"country_code,omitempty"`

	// An email to be used on the Draft Order.
	Email string `json:"email,omitempty"`

	BillingAddress Address `json:"billing_address,omitempty"`

	ShippingAddress Address `json:"shipping_address,omitempty"`

	// An array of Discount codes to add to the Draft Order.
	Discounts []PostDraftOrdersDraftOrderRequestDiscountsInner `json:"discounts,omitempty"`

	// An optional flag passed to the resulting order to determine use of notifications.
	NoNotificationOrder bool `json:"no_notification_order,omitempty"`

	// The ID of the Customer to associate the Draft Order with.
	CustomerId string `json:"customer_id,omitempty"`
}
