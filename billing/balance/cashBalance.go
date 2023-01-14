package customer

import (
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/intent"
)

type EuBankTransfer struct {
	// The BIC of the bank of the sender of the funding.
	Bic string `json:"bic,omitempty"`
	// The last 4 digits of the IBAN of the sender of the funding.
	IbanLast4 string `json:"iban_last4,omitempty"`
	// The full name of the sender, as supplied by the sending bank.
	SenderName string `json:"sender_name,omitempty"`
}


type BankTransfer struct {
	EuBankTransfer *EuBankTransfer `json:"eu_bank_transfer,omitempty"`
	// The user-supplied reference field on the bank transfer.
	Reference string `json:"reference,omitempty"`
	// The funding method type used to fund the customer balance. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
	Type string `json:"type"`
}


// CustomerCashBalanceTransaction Customers with certain payments enabled have a cash balance, representing funds that were paid by the customer to a merchant, but have not yet been allocated to a payment. Cash Balance Transactions represent when funds are moved into or out of this balance. This includes funding by the customer, allocation to payments, and refunds to the customer.
type CustomerCashBalanceTransaction struct {
	AppliedToPayment *intent.PaymentIntent `json:"applied_to_payment,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string                                 `json:"currency"`
	Customer *customer.Customer `json:"customer"`
	// The total available cash balance for the specified currency after this transaction was applied. Represented in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	EndingBalance int                                                                   `json:"ending_balance"`
	Funded        *BankTransfer `json:"funded,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The amount by which the cash balance changed, represented in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal). A positive value represents funds being added to the cash balance, a negative value represents funds being removed from the cash balance.
	NetAmount int `json:"net_amount"`
	// String representing the object's type. Objects of the same type share the same value.
	Object              string                                                                               `json:"object"`
	// RefundedFromPayment *refund.Refund `json:"refunded_from_payment,omitempty"`
	// The type of the cash balance transaction. One of `applied_to_payment`, `unapplied_from_payment`, `refunded_from_payment`, `funded`, `return_initiated`, or `return_canceled`. New types may be added in future. See [Customer Balance](https://stripe.com/docs/payments/customer-balance#types) to learn more about these types.
	Type                 string                                                                                `json:"type"`
	UnappliedFromPayment *intent.PaymentIntent `json:"unapplied_from_payment,omitempty"`
}
