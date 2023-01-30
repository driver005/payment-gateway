package methods

import "github.com/driver005/gateway/core"

// MandateBacsDebit
type MandateBacsDebit struct {
	core.Model

	// The status of the mandate on the Bacs network. Can be one of `pending`, `revoked`, `refused`, or `accepted`.
	NetworkStatus string `json:"network_status,omitempty"`
	// The unique reference identifying the mandate on the Bacs network.
	Reference string `json:"reference,omitempty"`
	// The URL that will contain the mandate that the customer has signed.
	Url string `json:"url,omitempty"`
}
