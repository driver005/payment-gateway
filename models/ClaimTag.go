package models

import "github.com/google/uuid"

// ClaimTag - Claim Tags are user defined tags that can be assigned to claim items for easy filtering and grouping.
type ClaimTag struct {
	Id uuid.UUID `json:"id database:primarykey"`

	// The value that the claim tag holds
	Value string `json:"value"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
