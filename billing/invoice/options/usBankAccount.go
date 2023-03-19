package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptionsUsBankAccount
type InvoicePaymentMethodOptionsUsBankAccount struct {
	core.Model

	FinancialConnections *InvoicePaymentMethodOptionsUsBankAccountLinkedAccountOptions `json:"financial_connections,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
