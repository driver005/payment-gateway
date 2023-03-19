package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptions
type InvoicePaymentMethodOptions struct {
	core.Model

	AcssDebit       *InvoicePaymentMethodOptionsAcssDebit       `json:"acss_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Bancontact      *InvoicePaymentMethodOptionsBancontact      `json:"bancontact,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Card            *InvoicePaymentMethodOptionsCard            `json:"card,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CustomerBalance *InvoicePaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Konbini         *InvoicePaymentMethodOptionsKonbini         `json:"konbini,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UsBankAccount   *InvoicePaymentMethodOptionsUsBankAccount   `json:"us_bank_account,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
