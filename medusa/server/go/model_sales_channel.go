/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// SalesChannel - A Sales Channel
type SalesChannel struct {

	// The sales channel's ID
	Id string `json:"id,omitempty"`

	// The name of the sales channel.
	Name string `json:"name"`

	// The description of the sales channel.
	Description string `json:"description,omitempty"`

	// Specify if the sales channel is enabled or disabled.
	IsDisabled bool `json:"is_disabled,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}