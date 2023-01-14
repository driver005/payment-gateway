package options

// InvoicePaymentMethodOptionsCustomerBalance
type InvoicePaymentMethodOptionsCustomerBalance struct {
	BankTransfer *InvoicePaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType string `json:"funding_type,omitempty"`
}
