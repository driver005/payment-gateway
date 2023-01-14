package models

// Networks
type Networks struct {
	// All available networks for the card.
	Available []string `json:"available"`
	// The preferred network for the card.
	Preferred string `json:"preferred,omitempty"`
}
