package fee

// FeeRefund `Application Fee Refund` objects allow you to refund an application fee that has previously been created but not yet refunded. Funds will be refunded to the Stripe account from which the fee was originally collected.  Related guide: [Refunding Application Fees](https://stripe.com/docs/connect/destination-charges#refunding-app-fee).
type FeeRefund struct {
	// Amount, in %s.
	Amount int `json:"amount"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
