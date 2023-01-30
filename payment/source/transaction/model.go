package transaction

import (
	"github.com/driver005/gateway/core"
)

// SourceTransaction Some payment methods have no required amount that a customer must send. Customers can be instructed to send any amount, and it can be made up of multiple transactions. As such, sources can have multiple associated transactions.
type SourceTransaction struct {
	core.Model

	AchCreditTransfer *SourceTransactionAchCreditTransferData `json:"ach_credit_transfer,omitempty" database:"foreignKey:id"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount your customer has pushed to the receiver.
	Amount            int                                     `json:"amount,omitempty"`
	ChfCreditTransfer *SourceTransactionChfCreditTransferData `json:"chf_credit_transfer,omitempty" database:"foreignKey:id"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency           string                                   `json:"currency,omitempty"`
	GbpCreditTransfer  *SourceTransactionGbpCreditTransferData  `json:"gbp_credit_transfer,omitempty" database:"foreignKey:id"`
	PaperCheck         *SourceTransactionPaperCheckData         `json:"paper_check,omitempty" database:"foreignKey:id"`
	SepaCreditTransfer *SourceTransactionSepaCreditTransferData `json:"sepa_credit_transfer,omitempty" database:"foreignKey:id"`
	// The ID of the source this transaction is attached to.
	Source string `json:"source,omitempty"`
	// The status of the transaction, one of `succeeded`, `pending`, or `failed`.
	Status string `json:"status,omitempty"`
	// The type of source this transaction is attached to.
	Type string `json:"type,omitempty"`
}
