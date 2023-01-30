package account

// AccountBacsDebitPaymentsSettings
type AccountBacsDebitPaymentsSettings struct {
	// The Bacs Direct Debit Display Name for this account. For payments made with Bacs Direct Debit, this will appear on the mandate, and as the statement descriptor.
	DisplayName string `json:"display_name,omitempty"`
}
