package options

// InvoicePaymentMethodOptionsCustomerBalanceBankTransfer
type InvoicePaymentMethodOptionsCustomerBalanceBankTransfer struct {
	EuBankTransfer *InvoicePaymentMethodOptionsCustomerBalanceBankTransferEuBankTransfer `json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
	Type string `json:"type,omitempty"`
}
