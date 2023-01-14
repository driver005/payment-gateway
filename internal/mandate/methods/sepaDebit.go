package methods

// MandateSepaDebit
type MandateSepaDebit struct {
	// The unique reference of the mandate.
	Reference string `json:"reference"`
	// The URL of the mandate. This URL generally contains sensitive information about the customer and should be shared with them exclusively.
	Url string `json:"url"`
}
