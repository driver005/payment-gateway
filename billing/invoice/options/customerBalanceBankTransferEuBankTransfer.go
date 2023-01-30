package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptionsCustomerBalanceBankTransferEuBankTransfer
type InvoicePaymentMethodOptionsCustomerBalanceBankTransferEuBankTransfer struct {
	core.Model

	// The desired country code of the bank account information. Permitted values include: `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country string `json:"country,omitempty"`
}
