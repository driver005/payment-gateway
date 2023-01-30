package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptionsBancontact
type InvoicePaymentMethodOptionsBancontact struct {
	core.Model

	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage string `json:"preferred_language,omitempty"`
}
