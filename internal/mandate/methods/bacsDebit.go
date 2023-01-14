package methods

// MandateBacsDebit
type MandateBacsDebit struct {
	// The status of the mandate on the Bacs network. Can be one of `pending`, `revoked`, `refused`, or `accepted`.
	NetworkStatus string `json:"network_status"`
	// The unique reference identifying the mandate on the Bacs network.
	Reference string `json:"reference"`
	// The URL that will contain the mandate that the customer has signed.
	Url string `json:"url"`
}
