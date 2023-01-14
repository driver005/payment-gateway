package details

// PaymentMethodDetailsAchDebit
type PaymentMethodDetailsAchDebit struct {
	// Type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType string `json:"account_holder_type,omitempty"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name,omitempty"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country,omitempty"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4,omitempty"`
	// Routing transit number of the bank account.
	RoutingNumber string `json:"routing_number,omitempty"`
}
