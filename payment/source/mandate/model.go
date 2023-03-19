package mandate

import (
	"github.com/driver005/gateway/core"
)

// SourceMandateNotificationSepaDebitData
type SourceMandateNotificationSepaDebitData struct {
	core.Model

	// SEPA creditor ID.
	CreditorIdentifier string `json:"creditor_identifier,omitempty"`
	// Last 4 digits of the account number associated with the debit.
	Last4 string `json:"last4,omitempty"`
	// Mandate reference associated with the debit.
	MandateReference string `json:"mandate_reference,omitempty"`
}

// SourceMandateNotification Source mandate notifications should be created when a notification related to a source mandate must be sent to the payer. They will trigger a webhook or deliver an email to the customer.
type SourceMandateNotification struct {
	core.Model

	AcssDebitStatementDescriptor string `json:"statement_descriptor_acss_debit,omitempty"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount associated with the mandate notification. The amount is expressed in the currency of the underlying source. Required if the notification type is `debit_initiated`.
	Amount         int    `json:"amount,omitempty"`
	BacsDebitLast4 string `json:"bacs_debit_last4,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	// The reason of the mandate notification. Valid reasons are `mandate_confirmed` or `debit_initiated`.
	Reason    string                                  `json:"reason,omitempty"`
	SepaDebit *SourceMandateNotificationSepaDebitData `json:"sepa_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Source    source.Source                           `json:"source,omitempty" database:"foreignKey:id"`
	// The status of the mandate notification. Valid statuses are `pending` or `submitted`.
	Status string `json:"status,omitempty"`
	// The type of source this mandate notification is attached to. Should be the source type identifier code for the payment method, such as `three_d_secure`.
	Type string `json:"type,omitempty"`
}
