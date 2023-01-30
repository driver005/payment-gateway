package page

import "github.com/driver005/gateway/core"

// PaymentPagesCheckoutSessionConsent
type PaymentPagesCheckoutSessionConsent struct {
	core.Model

	// If `opt_in`, the customer consents to receiving promotional communications from the merchant about this Checkout Session.
	Promotions string `json:"promotions,omitempty"`
	// If `accepted`, the customer in this Checkout Session has agreed to the merchant's terms of service.
	TermsOfService string `json:"terms_of_service,omitempty"`
}
