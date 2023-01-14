package methods

// PaymentMethodSepaDebit
type PaymentMethodSepaDebit struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code,omitempty"`
	// Branch code of bank associated with the bank account.
	BranchCode string `json:"branch_code,omitempty"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country,omitempty"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint   string                              `json:"fingerprint,omitempty"`
	GeneratedFrom PaymentMethodSepaDebitGeneratedFrom `json:"generated_from,omitempty"`
	// Last four characters of the IBAN.
	Last4 string `json:"last4,omitempty"`
}
