package bank

import (
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/models"
)

// BankAccount These bank accounts are payment methods on `Customer` objects.  On the other hand [External Accounts](https://stripe.com/docs/api#external_accounts) are transfer destinations on `Account` objects for [Custom accounts](https://stripe.com/docs/connect/custom-accounts). They can be bank accounts or debit cards as well, and are documented in the links above.  Related guide: [Bank Debits and Transfers](https://stripe.com/docs/payments/bank-debits-transfers).
type BankAccount struct {
	Account models.Account `json:"account,omitempty"`
	// The name of the person or business that owns the bank account.
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType string `json:"account_holder_type,omitempty"`
	// The bank account type. This can only be `checking` or `savings` in most countries. In Japan, this can only be `futsu` or `toza`.
	AccountType string `json:"account_type,omitempty"`
	// A set of available payout methods for this bank account. Only values from this set should be passed as the `method` when creating a payout.
	AvailablePayoutMethods []string `json:"available_payout_methods,omitempty"`
	// Name of the bank associated with the routing number (e.g., `WELLS FARGO`).
	BankName string `json:"bank_name,omitempty"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// Three-letter [ISO code for the currency](https://stripe.com/docs/payouts) paid out to the bank account.
	Currency string `json:"currency"`
	Customer customer.Customer `json:"customer,omitempty"`
	// Whether this bank account is the default external account for its currency.
	DefaultForCurrency bool `json:"default_for_currency,omitempty"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// The last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The routing transit number for the bank account.
	RoutingNumber string `json:"routing_number,omitempty"`
	// For bank accounts, possible values are `new`, `validated`, `verified`, `verification_failed`, or `errored`. A bank account that hasn't had any activity or validation performed is `new`. If Stripe can determine that the bank account exists, its status will be `validated`. Note that there often isn’t enough information to know (e.g., for smaller credit unions), and the validation is not always run. If customer bank account verification has succeeded, the bank account status will be `verified`. If the verification failed for any reason, such as microdeposit failure, the status will be `verification_failed`. If a transfer sent to this bank account fails, we'll set the status to `errored` and will not continue to send transfers until the bank details are updated.  For external accounts, possible values are `new` and `errored`. Validations aren't run against external accounts because they're only used for payouts. This means the other statuses don't apply. If a transfer fails, the status is set to `errored` and transfers are stopped until account details are updated.
	Status string `json:"status"`
}
