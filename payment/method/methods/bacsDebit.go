package methods

// PaymentMethodBacsDebit
type PaymentMethodBacsDebit struct {
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4,omitempty"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode string `json:"sort_code,omitempty"`
}
