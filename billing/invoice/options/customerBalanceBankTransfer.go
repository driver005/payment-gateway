package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptionsCustomerBalanceBankTransfer
type InvoicePaymentMethodOptionsCustomerBalanceBankTransfer struct {
	core.Model

	EuBankTransfer *InvoicePaymentMethodOptionsCustomerBalanceBankTransferEuBankTransfer `json:"eu_bank_transfer,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// The bank transfer type that can be used for funding. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
	Type string `json:"type,omitempty"`
}
