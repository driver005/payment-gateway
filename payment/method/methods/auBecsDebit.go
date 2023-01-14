package methods

// PaymentMethodAuBecsDebit
type PaymentMethodAuBecsDebit struct {
	// Six-digit number identifying bank and branch associated with this bank account.
	BsbNumber string `json:"bsb_number,omitempty"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4,omitempty"`
}
