package models

import "github.com/driver005/gateway/core"

// User - Represents a User who can manage store settings.
type User struct {
	core.Model

	// The email of the User
	Email string `json:"email"`

	// The first name of the User
	FirstName string `json:"first_name" database:"default:null"`

	// The last name of the User
	LastName string `json:"last_name" database:"default:null"`

	// An API token associated with the user.
	ApiToken string `json:"api_token" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
