package models

import "github.com/driver005/gateway/core"

// Images holds a reference to a URL at which the image file can be found.
type Image struct {
	core.Model

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`

	// The URL at which the image file can be found.
	Url string `json:"url"`
}
