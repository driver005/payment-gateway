package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptionsAcssDebit
type InvoicePaymentMethodOptionsAcssDebit struct {
	core.Model

	MandateOptionsTransactionType string `json:"mandate_options_transaction_type,omitempty"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
