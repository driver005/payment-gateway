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

// IdempotencyKey - Idempotency Key is used to continue a process in case of any failure that might occur.
type IdempotencyKey struct {

	// The idempotency key's ID
	Id string `json:"id,omitempty"`

	// The unique randomly generated key used to determine the state of a process.
	IdempotencyKey string `json:"idempotency_key"`

	// Date which the idempotency key was locked.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Date which the idempotency key was locked.
	LockedAt time.Time `json:"locked_at,omitempty"`

	// The method of the request
	RequestMethod string `json:"request_method,omitempty"`

	// The parameters passed to the request
	RequestParams map[string]interface{} `json:"request_params,omitempty"`

	// The request's path
	RequestPath string `json:"request_path,omitempty"`

	// The response's code.
	ResponseCode string `json:"response_code,omitempty"`

	// The response's body
	ResponseBody map[string]interface{} `json:"response_body,omitempty"`

	// Where to continue from.
	RecoveryPoint string `json:"recovery_point,omitempty"`
}
