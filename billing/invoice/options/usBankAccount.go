package options

// InvoicePaymentMethodOptionsUsBankAccount
type InvoicePaymentMethodOptionsUsBankAccount struct {
	FinancialConnections *InvoicePaymentMethodOptionsUsBankAccountLinkedAccountOptions `json:"financial_connections,omitempty"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
