package transaction

import "github.com/driver005/gateway/core"

// SourceTransactionGbpCreditTransferData
type SourceTransactionGbpCreditTransferData struct {
	core.Model

	// Bank account fingerprint associated with the Stripe owned bank account receiving the transfer.
	Fingerprint string `json:"fingerprint,omitempty"`
	// The credit transfer rails the sender used to push this transfer. The possible rails are: Faster Payments, BACS, CHAPS, and wire transfers. Currently only Faster Payments is supported.
	FundingMethod string `json:"funding_method,omitempty"`
	// Last 4 digits of sender account number associated with the transfer.
	Last4 string `json:"last4,omitempty"`
	// Sender entered arbitrary information about the transfer.
	Reference string `json:"reference,omitempty"`
	// Sender account number associated with the transfer.
	SenderAccountNumber string `json:"sender_account_number,omitempty"`
	// Sender name associated with the transfer.
	SenderName string `json:"sender_name,omitempty"`
	// Sender sort code associated with the transfer.
	SenderSortCode string `json:"sender_sort_code,omitempty"`
}
