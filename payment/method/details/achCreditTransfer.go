package details

// PaymentMethodDetailsAchCreditTransfer
type PaymentMethodDetailsAchCreditTransfer struct {
	// Account number to transfer funds to.
	AccountNumber string `json:"account_number,omitempty"`
	// Name of the bank associated with the routing number.
	BankName string `json:"bank_name,omitempty"`
	// Routing transit number for the bank account to transfer funds to.
	RoutingNumber string `json:"routing_number,omitempty"`
	// SWIFT code of the bank associated with the routing number.
	SwiftCode string `json:"swift_code,omitempty"`
}
