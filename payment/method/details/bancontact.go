package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsBancontact
type PaymentMethodDetailsBancontact struct {
	core.Model

	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code,omitempty"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name,omitempty"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic,omitempty"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4,omitempty"`
	// Preferred language of the Bancontact authorization page that the customer is redirected to. Can be one of `en`, `de`, `fr`, or `nl`
	PreferredLanguage string `json:"preferred_language,omitempty"`
	// Owner's verified full name. Values are verified or provided by Bancontact directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name,omitempty"`
}
