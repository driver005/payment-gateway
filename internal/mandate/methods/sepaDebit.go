package methods

import "github.com/driver005/gateway/core"

// MandateSepaDebit
type MandateSepaDebit struct {
	core.Model

	// The unique reference of the mandate.
	Reference string `json:"reference,omitempty"`
	// The URL of the mandate. This URL generally contains sensitive information about the customer and should be shared with them exclusively.
	Url string `json:"url,omitempty"`
}
