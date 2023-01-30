package cash

import (
	"github.com/driver005/gateway/core"
)

// CashBalance A customer's `Cash balance` represents real funds. Customers can add funds to their cash balance by sending a bank transfer. These funds can be used for payment and can eventually be paid out to your bank account.
type CashBalance struct {
	core.Model

	// A hash of all cash balances available to this customer. You cannot delete a customer with any cash balances, even if the balance is 0. Amounts are represented in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Available core.JSONB `json:"available,omitempty"`
	// The ID of the customer whose cash balance this object represents.
	Customer string `json:"customer,omitempty"`
	// The configuration for how funds that land in the customer cash balance are reconciled.
	SettingReconciliationMode string `json:"setting_reconciliation_mode,omitempty"`
	// A flag to indicate if reconciliation mode returned is the user's default or is specific to this customer cash balance
	SettingUsingMerchantDefault bool `json:"setting_using_merchant_default,omitempty"`
}
