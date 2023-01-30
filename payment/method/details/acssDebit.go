package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsAcssDebit
type PaymentMethodDetailsAcssDebit struct {
	core.Model

	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name,omitempty"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Institution number of the bank account
	InstitutionNumber string `json:"institution_number,omitempty"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4,omitempty"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate,omitempty"`
	// Transit number of the bank account.
	TransitNumber string `json:"transit_number,omitempty"`
}
