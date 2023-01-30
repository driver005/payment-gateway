package transaction

import "github.com/driver005/gateway/core"

// SourceTransactionSepaCreditTransferData
type SourceTransactionSepaCreditTransferData struct {
	core.Model

	// Reference associated with the transfer.
	Reference string `json:"reference,omitempty"`
	// Sender's bank account IBAN.
	SenderIban string `json:"sender_iban,omitempty"`
	// Sender's name.
	SenderName string `json:"sender_name,omitempty"`
}
