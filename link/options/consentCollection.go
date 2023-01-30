package options

import "github.com/driver005/gateway/core"

// PaymentLinksResourceConsentCollection
type PaymentLinksResourceConsentCollection struct {
	core.Model

	// If set to `auto`, enables the collection of customer consent for promotional communications.
	Promotions string `json:"promotions,omitempty"`
	// If set to `required`, it requires cutomers to accept the terms of service before being able to pay. If set to `none`, customers won't be shown a checkbox to accept the terms of service.
	TermsOfService string `json:"terms_of_service,omitempty"`
}
