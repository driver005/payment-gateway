package methods

// PaymentMethodUsBankAccount
type PaymentMethodUsBankAccount struct {
	// Account holder type: individual or company.
	AccountHolderType string `json:"account_holder_type,omitempty"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType string `json:"account_type,omitempty"`
	// The name of the bank.
	BankName string `json:"bank_name,omitempty"`
	// The ID of the Financial Connections Account used to create the payment method.
	FinancialConnectionsAccount string `json:"financial_connections_account,omitempty"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Last four digits of the bank account number.
	Last4    string                             `json:"last4,omitempty"`
	Networks PaymentMethodUsBankAccountNetworks `json:"networks,omitempty"`
	// Routing number of the bank account.
	RoutingNumber string `json:"routing_number,omitempty"`
}
