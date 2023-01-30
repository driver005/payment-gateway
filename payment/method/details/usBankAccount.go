package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsUsBankAccount
type PaymentMethodDetailsUsBankAccount struct {
	core.Model

	// Account holder type: individual or company.
	AccountHolderType string `json:"account_holder_type,omitempty"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType string `json:"account_type,omitempty"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name,omitempty"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4,omitempty"`
	// Routing number of the bank account.
	RoutingNumber string `json:"routing_number,omitempty"`
}