package balance

import "github.com/driver005/gateway/core"

// BalanceDetail
type BalanceDetail struct {
	core.Model

	// Funds that are available for use.
	Available []BalanceAmount `json:"available,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
}

// BalanceAmountBySourceType
type BalanceAmountBySourceType struct {
	core.Model

	// Amount for bank account.
	BankAccount int `json:"bank_account,omitempty"`
	// Amount for card.
	Card int `json:"card,omitempty"`
	// Amount for FPX.
	Fpx int `json:"fpx,omitempty"`
}

// BalanceAmount
type BalanceAmount struct {
	core.Model

	// Balance amount.
	Amount int `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency    string                     `json:"currency,omitempty"`
	SourceTypes *BalanceAmountBySourceType `json:"source_types,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}

// Balance This is an object representing your Stripe balance. You can retrieve it to see the balance currently on your Stripe account.  You can also retrieve the balance history, which contains a list of [transactions](https://stripe.com/docs/reporting/balance-transaction-types) that contributed to the balance (charges, payouts, and so forth).  The available and pending amounts for each currency are broken down further by payment source types.  Related guide: [Understanding Connect Account Balances](https://stripe.com/docs/connect/account-balances).
type Balance struct {
	core.Model

	// Funds that are available to be transferred or paid out, whether automatically by Stripe or explicitly via the [Transfers API](https://stripe.com/docs/api#transfers) or [Payouts API](https://stripe.com/docs/api#payouts). The available balance for each currency and payment type can be found in the `source_types` property.
	Available []BalanceAmount `json:"available,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	// Funds held due to negative balances on connected Custom accounts. The connect reserve balance for each currency and payment type can be found in the `source_types` property.
	ConnectReserved []BalanceAmount `json:"connect_reserved,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	// Funds that can be paid out using Instant Payouts.
	InstantAvailable []BalanceAmount `json:"instant_available,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	Issuing          *BalanceDetail  `json:"issuing,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Funds that are not yet available in the balance, due to the 7-day rolling pay cycle. The pending balance for each currency, and for each payment type, can be found in the `source_types` property.
	Pending []BalanceAmount `json:"pending,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
}
