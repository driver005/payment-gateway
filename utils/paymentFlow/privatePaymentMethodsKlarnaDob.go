package paymentFlow

import "github.com/driver005/gateway/core"

// PaymentFlowsPrivatePaymentMethodsKlarnaDob
type PaymentFlowsPrivatePaymentMethodsKlarnaDob struct {
	core.Model

	// The day of birth, between 1 and 31.
	Day int `json:"day,omitempty"`
	// The month of birth, between 1 and 12.
	Month int `json:"month,omitempty"`
	// The four-digit year of birth.
	Year int `json:"year,omitempty"`
}
