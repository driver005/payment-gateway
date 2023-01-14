package attempt

import (
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/payment/method"
)

// SetupAttemptPaymentMethodDetailsSofort
type SetupAttemptPaymentMethodDetailsSofort struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code,omitempty"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name,omitempty"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic                       string               `json:"bic,omitempty"`
	GeneratedSepaDebit        method.PaymentMethod `json:"generated_sepa_debit,omitempty"`
	GeneratedSepaDebitMandate mandate.Mandate      `json:"generated_sepa_debit_mandate,omitempty"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4,omitempty"`
	// Preferred language of the Sofort authorization page that the customer is redirected to. Can be one of `en`, `de`, `fr`, or `nl`
	PreferredLanguage string `json:"preferred_language,omitempty"`
	// Owner's verified full name. Values are verified or provided by Sofort directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name,omitempty"`
}
