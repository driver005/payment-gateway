package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsCardChecks
type PaymentMethodDetailsCardChecks struct {
	core.Model

	// If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check string `json:"address_line1_check,omitempty"`
	// If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressPostalCodeCheck string `json:"address_postal_code_check,omitempty"`
	// If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	CvcCheck string `json:"cvc_check,omitempty"`
}
