package transaction

// SourceTransaction Some payment methods have no required amount that a customer must send. Customers can be instructed to send any amount, and it can be made up of multiple transactions. As such, sources can have multiple associated transactions.
type SourceTransaction struct {
	AchCreditTransfer *SourceTransactionAchCreditTransferData `json:"ach_credit_transfer,omitempty"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount your customer has pushed to the receiver.
	Amount int `json:"amount"`
	ChfCreditTransfer *SourceTransactionChfCreditTransferData `json:"chf_credit_transfer,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency"`
	GbpCreditTransfer *SourceTransactionGbpCreditTransferData `json:"gbp_credit_transfer,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	PaperCheck *SourceTransactionPaperCheckData `json:"paper_check,omitempty"`
	SepaCreditTransfer *SourceTransactionSepaCreditTransferData `json:"sepa_credit_transfer,omitempty"`
	// The ID of the source this transaction is attached to.
	Source string `json:"source"`
	// The status of the transaction, one of `succeeded`, `pending`, or `failed`.
	Status string `json:"status"`
	// The type of source this transaction is attached to.
	Type string `json:"type"`
}
