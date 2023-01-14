package transaction

// SourceTransactionAchCreditTransferData
type SourceTransactionAchCreditTransferData struct {
	// Customer data associated with the transfer.
	CustomerData *string `json:"customer_data,omitempty"`
	// Bank account fingerprint associated with the transfer.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Last 4 digits of the account number associated with the transfer.
	Last4 *string `json:"last4,omitempty"`
	// Routing number associated with the transfer.
	RoutingNumber *string `json:"routing_number,omitempty"`
}
