package transaction

// SourceTransactionChfCreditTransferData
type SourceTransactionChfCreditTransferData struct {
	// Reference associated with the transfer.
	Reference *string `json:"reference,omitempty"`
	// Sender's country address.
	SenderAddressCountry *string `json:"sender_address_country,omitempty"`
	// Sender's line 1 address.
	SenderAddressLine1 *string `json:"sender_address_line1,omitempty"`
	// Sender's bank account IBAN.
	SenderIban *string `json:"sender_iban,omitempty"`
	// Sender's name.
	SenderName *string `json:"sender_name,omitempty"`
}
