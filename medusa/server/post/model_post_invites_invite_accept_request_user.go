/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PostInvitesInviteAcceptRequestUser - The User to create.
type PostInvitesInviteAcceptRequestUser struct {

	// the first name of the User
	FirstName string `json:"first_name"`

	// the last name of the User
	LastName string `json:"last_name"`

	// The desired password for the User
	Password string `json:"password"`
}
