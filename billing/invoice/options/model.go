package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptions
type InvoicePaymentMethodOptions struct {
	core.Model

	AcssDebit       *InvoicePaymentMethodOptionsAcssDebit       `json:"acss_debit,omitempty" database:"foreignKey:id"`
	Bancontact      *InvoicePaymentMethodOptionsBancontact      `json:"bancontact,omitempty" database:"foreignKey:id"`
	Card            *InvoicePaymentMethodOptionsCard            `json:"card,omitempty" database:"foreignKey:id"`
	CustomerBalance *InvoicePaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty" database:"foreignKey:id"`
	Konbini         *InvoicePaymentMethodOptionsKonbini         `json:"konbini,omitempty" database:"foreignKey:id"`
	UsBankAccount   *InvoicePaymentMethodOptionsUsBankAccount   `json:"us_bank_account,omitempty" database:"foreignKey:id"`
}
