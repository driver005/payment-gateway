package mandate

import "github.com/driver005/gateway/payment/source"

// SourceMandateNotification Source mandate notifications should be created when a notification related to a source mandate must be sent to the payer. They will trigger a webhook or deliver an email to the customer.
type SourceMandateNotification struct {
	AcssDebit *SourceMandateNotificationAcssDebitData `json:"acss_debit,omitempty"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount associated with the mandate notification. The amount is expressed in the currency of the underlying source. Required if the notification type is `debit_initiated`.
	Amount int `json:"amount,omitempty"`
	BacsDebit *SourceMandateNotificationBacsDebitData `json:"bacs_debit,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The reason of the mandate notification. Valid reasons are `mandate_confirmed` or `debit_initiated`.
	Reason string `json:"reason"`
	SepaDebit *SourceMandateNotificationSepaDebitData `json:"sepa_debit,omitempty"`
	Source source.Source `json:"source"`
	// The status of the mandate notification. Valid statuses are `pending` or `submitted`.
	Status string `json:"status"`
	// The type of source this mandate notification is attached to. Should be the source type identifier code for the payment method, such as `three_d_secure`.
	Type string `json:"type"`
}
