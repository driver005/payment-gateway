package details

// PaymentMethodDetailsGiropay
type PaymentMethodDetailsGiropay struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code,omitempty"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name,omitempty"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic,omitempty"`
	// Owner's verified full name. Values are verified or provided by Giropay directly (if supported) at the time of authorization or settlement. They cannot be set or mutated. Giropay rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name,omitempty"`
}
