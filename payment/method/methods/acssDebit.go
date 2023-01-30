package methods

import "github.com/driver005/gateway/core"

type PaymentMethodAcssDebit struct {
	core.Model

	BankName          string `json:"bank_name,omitempty"`
	Fingerprint       string `json:"fingerprint,omitempty"`
	InstitutionNumber string `json:"institution_number,omitempty"`
	Last4             string `json:"last4,omitempty"`
	TransitNumber     string `json:"transit_number,omitempty"`
}
