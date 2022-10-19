package models

type CustomerBalances struct {
	ID            string      `json:"id"`
	Object        string      `json:"object"`
	Amount        int         `json:"amount"`
	Created       int         `json:"created"`
	CreditNote    interface{} `json:"credit_note"`
	Currency      string      `json:"currency"`
	Customer      string      `json:"customer"`
	Description   interface{} `json:"description"`
	EndingBalance int         `json:"ending_balance"`
	Invoice       interface{} `json:"invoice"`
	Livemode      bool        `json:"livemode"`
	Metadata      interface{} `json:"metadata"`
	Type          string      `json:"type"`
}
