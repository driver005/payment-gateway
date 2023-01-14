package options

// SubscriptionsResourcePaymentMethodOptions
type SubscriptionsResourcePaymentMethodOptions struct {
	AcssDebit       SubscriptionsResourcePaymentMethodOptionsAcssDebit       `json:"acss_debit,omitempty"`
	Bancontact      SubscriptionsResourcePaymentMethodOptionsBancontact      `json:"bancontact,omitempty"`
	Card            SubscriptionsResourcePaymentMethodOptionsCard            `json:"card,omitempty"`
	CustomerBalance SubscriptionsResourcePaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty"`
	Konbini         SubscriptionsResourcePaymentMethodOptionsKonbini         `json:"konbini,omitempty"`
	UsBankAccount   SubscriptionsResourcePaymentMethodOptionsUsBankAccount   `json:"us_bank_account,omitempty"`
}
