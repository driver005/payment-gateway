package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptionsCustomerBalance
type InvoicePaymentMethodOptionsCustomerBalance struct {
	core.Model

	BankTransfer *InvoicePaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer,omitempty" database:"foreignKey:id"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType string `json:"funding_type,omitempty"`
}
