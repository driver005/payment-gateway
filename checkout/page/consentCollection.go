package page

import "github.com/driver005/gateway/core"

// PaymentPagesCheckoutSessionConsentCollection
type PaymentPagesCheckoutSessionConsentCollection struct {
	core.Model

	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout Session will determine whether to display an option to opt into promotional communication from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions string `json:"promotions,omitempty"`
	// If set to `required`, it requires customers to accept the terms of service before being able to pay.
	TermsOfService string `json:"terms_of_service,omitempty"`
}
