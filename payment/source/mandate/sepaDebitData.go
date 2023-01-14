package mandate

// SourceMandateNotificationSepaDebitData
type SourceMandateNotificationSepaDebitData struct {
	// SEPA creditor ID.
	CreditorIdentifier *string `json:"creditor_identifier,omitempty"`
	// Last 4 digits of the account number associated with the debit.
	Last4 *string `json:"last4,omitempty"`
	// Mandate reference associated with the debit.
	MandateReference *string `json:"mandate_reference,omitempty"`
}