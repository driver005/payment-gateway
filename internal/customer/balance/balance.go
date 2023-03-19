package customer

import (
	"github.com/driver005/gateway/internal/customer"
	"github.com/google/uuid"
)

// CustomerBalanceCustomerBalanceSettings
type CustomerBalanceCustomerBalanceSettings struct {
	// The configuration for how funds that land in the customer cash balance are reconciled.
	ReconciliationMode string `json:"reconciliation_mode,omitempty"`
	// A flag to indicate if reconciliation mode returned is the user's default or is specific to this customer cash balance
	UsingMerchantDefault bool `json:"using_merchant_default,omitempty"`
}

// CustomerBalanceTransaction Each customer has a [`balance`](https://stripe.com/docs/api/customers/object#customer_object-balance) value, which denotes a debit or credit that's automatically applied to their next invoice upon finalization. You may modify the value directly by using the [update customer API](https://stripe.com/docs/api/customers/update), or by creating a Customer Balance Transaction, which increments or decrements the customer's `balance` by the specified `amount`.  Related guide: [Customer Balance](https://stripe.com/docs/billing/customer/balance) to learn more.
type CustomerBalanceTransaction struct {
	// The amount of the transaction. A negative value is a credit for the customer's balance, and a positive value is a debit to the customer's `balance`.
	Amount int `json:"amount,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	// CreditNote *CreditNote `json:"credit_note,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`

	CustomerId *uuid.UUID         `json:"customer_id,omitempty" swaggerignore:"true"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id" swaggertype:"primitive,string" format:"uuid"`

	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// The customer's `balance` after the transaction was applied. A negative value decreases the amount due on the customer's next invoice. A positive value increases the amount due on the customer's next invoice.
	EndingBalance int `json:"ending_balance,omitempty"`
	// Transaction type: `adjustment`, `applied_to_invoice`, `credit_note`, `initial`, `invoice_overpaid`, `invoice_too_large`, `invoice_too_small`, `unspent_receiver_credit`, or `unapplied_from_invoice`. See the [Customer Balance page](https://stripe.com/docs/billing/customer/balance#types) to learn more about transaction types.
	Type Type `json:"type,omitempty"`
}
