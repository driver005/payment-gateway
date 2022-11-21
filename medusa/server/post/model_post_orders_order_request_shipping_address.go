/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PostOrdersOrderRequestShippingAddress - Shipping address
type PostOrdersOrderRequestShippingAddress struct {

	// ID of the address
	Id string `json:"id,omitempty"`

	// ID of the customer this address belongs to
	CustomerId string `json:"customer_id,omitempty"`

	// Available if the relation `customer` is expanded.
	Customer []map[string]interface{} `json:"customer,omitempty"`

	// Company name
	Company string `json:"company,omitempty"`

	// First name
	FirstName string `json:"first_name,omitempty"`

	// Last name
	LastName string `json:"last_name,omitempty"`

	// Address line 1
	Address1 string `json:"address_1,omitempty"`

	// Address line 2
	Address2 string `json:"address_2,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// The 2 character ISO code of the country in lower case
	CountryCode string `json:"country_code,omitempty"`

	// A country object. Available if the relation `country` is expanded.
	Country map[string]interface{} `json:"country,omitempty"`

	// Province
	Province string `json:"province,omitempty"`

	// Postal Code
	PostalCode string `json:"postal_code,omitempty"`

	// Phone Number
	Phone string `json:"phone,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
