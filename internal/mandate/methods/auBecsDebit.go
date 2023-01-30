package methods

import "github.com/driver005/gateway/core"

// MandateAuBecsDebit
type MandateAuBecsDebit struct {
	core.Model

	// The URL of the mandate. This URL generally contains sensitive information about the customer and should be shared with them exclusively.
	Url string `json:"url,omitempty"`
}
