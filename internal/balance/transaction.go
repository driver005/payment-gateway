package balance

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/fee"
)

// BalanceTransaction Balance transactions represent funds moving through your Stripe account. They're created for every type of transaction that comes into or flows out of your Stripe account balance.  Related guide: [Balance Transaction Types](https://stripe.com/docs/reports/balance-transaction-types).
type BalanceTransaction struct {
	core.Model
	// Gross amount of the transaction, in %s.
	Amount int `json:"amount,omitempty"`
	// The date the transaction's net funds will become available in the Stripe balance.
	AvailableOn int `json:"available_on,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// The exchange rate used, if applicable, for this transaction. Specifically, if money was converted from currency A to currency B, then the `amount` in currency A, times `exchange_rate`, would be the `amount` in currency B. For example, suppose you charged a customer 10.00 EUR. Then the PaymentIntent's `amount` would be `1000` and `currency` would be `eur`. Suppose this was converted into 12.34 USD in your Stripe account. Then the BalanceTransaction's `amount` would be `1234`, `currency` would be `usd`, and `exchange_rate` would be `1.234`.
	ExchangeRate float64 `json:"exchange_rate,omitempty"`
	// Fees (in %s) paid for this transaction.
	Fee int `json:"fee,omitempty"`
	// Detailed breakdown of fees (in %s) paid for this transaction.
	FeeDetails []fee.Fee `json:"fee_details,omitempty" database:"foreignKey:id"`
	// Net amount of the transaction, in %s.
	Net int `json:"net,omitempty"`
	// [Learn more](https://stripe.com/docs/reports/reporting-categories) about how reporting categories can help you understand balance transactions from an accounting perspective.
	ReportingCategory string `json:"reporting_category,omitempty"`
	Source            string `json:"source,omitempty"`
	// If the transaction's net funds are available in the Stripe balance yet. Either `available` or `pending`.
	Status string `json:"status,omitempty"`
	// Transaction type: `adjustment`, `advance`, `advance_funding`, `anticipation_repayment`, `application_fee`, `application_fee_refund`, `charge`, `connect_collection_transfer`, `contribution`, `issuing_authorization_hold`, `issuing_authorization_release`, `issuing_dispute`, `issuing_transaction`, `payment`, `payment_failure_refund`, `payment_refund`, `payout`, `payout_cancel`, `payout_failure`, `refund`, `refund_failure`, `reserve_transaction`, `reserved_funds`, `stripe_fee`, `stripe_fx_fee`, `tax_fee`, `topup`, `topup_reversal`, `transfer`, `transfer_cancel`, `transfer_failure`, or `transfer_refund`. [Learn more](https://stripe.com/docs/reports/balance-transaction-types) about balance transaction types and what they represent. If you are looking to classify transactions for accounting purposes, you might want to consider `reporting_category` instead.
	Type string `json:"type,omitempty"`
}
