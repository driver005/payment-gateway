package models

type Refund struct {
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	Amount             int         `json:"amount"`
	BalanceTransaction interface{} `json:"balance_transaction"`
	Charge             string      `json:"charge"`
	Created            int         `json:"created"`
	Currency           string      `json:"currency"`
	Metadata           struct {
	} `json:"metadata"`
	PaymentIntent          interface{} `json:"payment_intent"`
	Reason                 interface{} `json:"reason"`
	ReceiptNumber          interface{} `json:"receipt_number"`
	SourceTransferReversal interface{} `json:"source_transfer_reversal"`
	Status                 string      `json:"status"`
	TransferReversal       interface{} `json:"transfer_reversal"`
}