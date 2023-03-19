package page

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
)

type TermsOfService string

const (
	TermsOfServiceAccepted TermsOfService = "accepted"
)

func (ct *TermsOfService) Scan(value interface{}) error {
	*ct = TermsOfService(value.([]byte))
	return nil
}

func (ct TermsOfService) Value() (driver.Value, error) {
	return string(ct), nil
}

// PaymentPagesCheckoutSessionConsent
type PaymentPagesCheckoutSessionConsent struct {
	core.Model

	// If `opt_in`, the customer consents to receiving promotional communications from the merchant about this Checkout Session.
	Promotions string `json:"promotions,omitempty"`
	// If `accepted`, the customer in this Checkout Session has agreed to the merchant's terms of service.
	TermsOfService TermsOfService `json:"terms_of_service,omitempty"`
}
