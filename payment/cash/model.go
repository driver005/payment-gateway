package cash

// BalanceSettings
type BalanceSettings struct {
	// The configuration for how funds that land in the customer cash balance are reconciled.
	ReconciliationMode string `json:"reconciliation_mode"`
	// A flag to indicate if reconciliation mode returned is the user's default or is specific to this customer cash balance
	UsingMerchantDefault bool `json:"using_merchant_default"`
}

// CashBalance A customer's `Cash balance` represents real funds. Customers can add funds to their cash balance by sending a bank transfer. These funds can be used for payment and can eventually be paid out to your bank account.
type CashBalance struct {
	// A hash of all cash balances available to this customer. You cannot delete a customer with any cash balances, even if the balance is 0. Amounts are represented in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Available map[string]int `json:"available,omitempty"`
	// The ID of the customer whose cash balance this object represents.
	Customer string `json:"customer"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	Settings BalanceSettings `json:"settings"`
}
