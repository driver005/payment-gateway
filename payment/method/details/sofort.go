package details

// PaymentMethodDetailsSofort
type PaymentMethodDetailsSofort struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code,omitempty"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name,omitempty"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic,omitempty"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country                   string                                                  `json:"country,omitempty"`
	GeneratedSepaDebit        PaymentMethodDetailsBancontactGeneratedSepaDebit        `json:"generated_sepa_debit,omitempty"`
	GeneratedSepaDebitMandate PaymentMethodDetailsBancontactGeneratedSepaDebitMandate `json:"generated_sepa_debit_mandate,omitempty"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4,omitempty"`
	// Preferred language of the SOFORT authorization page that the customer is redirected to. Can be one of `de`, `en`, `es`, `fr`, `it`, `nl`, or `pl`
	PreferredLanguage string `json:"preferred_language,omitempty"`
	// Owner's verified full name. Values are verified or provided by SOFORT directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name,omitempty"`
}
