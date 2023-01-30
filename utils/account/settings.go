package account

// AccountSettings
type AccountSettings struct {
	BacsDebitPayments *AccountBacsDebitPaymentsSettings `json:"bacs_debit_payments,omitempty"`
	Branding          AccountBrandingSettings           `json:"branding"`
	CardIssuing       *AccountCardIssuingSettings       `json:"card_issuing,omitempty"`
	CardPayments      AccountCardPaymentsSettings       `json:"card_payments"`
	Dashboard         AccountDashboardSettings          `json:"dashboard"`
	Payments          AccountPaymentsSettings           `json:"payments"`
	Payouts           *AccountPayoutSettings            `json:"payouts,omitempty"`
	SepaDebitPayments *AccountSepaDebitPaymentsSettings `json:"sepa_debit_payments,omitempty"`
	Treasury          *AccountTreasurySettings          `json:"treasury,omitempty"`
}
