package fee

import (
	"github.com/driver005/gateway/core"
)

// Fee
type Fee struct {
	core.Model

	// Amount of the fee, in cents.
	Amount int `json:"amount,omitempty"`
	// ID of the Connect application that earned the fee.
	Application string `json:"application,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// Type of the fee, one of: `application_fee`, `stripe_fee` or `tax`.
	Type string `json:"type,omitempty"`
}

// FeeRefund `Application Fee Refund` objects allow you to refund an application fee that has previously been created but not yet refunded. Funds will be refunded to the Stripe account from which the fee was originally collected.  Related guide: [Refunding Application Fees](https://stripe.com/docs/connect/destination-charges#refunding-app-fee).
type FeeRefund struct {
	core.Model

	// Amount, in %s.
	Amount int `json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
}
