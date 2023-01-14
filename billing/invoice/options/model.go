package options

// InvoicePaymentMethodOptions
type InvoicePaymentMethodOptions struct {
	AcssDebit       *InvoicePaymentMethodOptionsAcssDebit       `json:"acss_debit,omitempty"`
	Bancontact      *InvoicePaymentMethodOptionsBancontact      `json:"bancontact,omitempty"`
	Card            *InvoicePaymentMethodOptionsCard            `json:"card,omitempty"`
	CustomerBalance *InvoicePaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty"`
	Konbini         *InvoicePaymentMethodOptionsKonbini         `json:"konbini,omitempty"`
	UsBankAccount   *InvoicePaymentMethodOptionsUsBankAccount   `json:"us_bank_account,omitempty"`
}
