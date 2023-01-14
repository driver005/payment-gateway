package transaction

// SourceTransactionSepaCreditTransferData
type SourceTransactionSepaCreditTransferData struct {
	// Reference associated with the transfer.
	Reference *string `json:"reference,omitempty"`
	// Sender's bank account IBAN.
	SenderIban *string `json:"sender_iban,omitempty"`
	// Sender's name.
	SenderName *string `json:"sender_name,omitempty"`
}
