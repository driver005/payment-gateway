package models

type Payout struct {
	ID                        string      `json:"id"`
	Object                    string      `json:"object"`
	Amount                    int         `json:"amount"`
	ArrivalDate               int         `json:"arrival_date"`
	Automatic                 bool        `json:"automatic"`
	BalanceTransaction        string      `json:"balance_transaction"`
	Created                   int         `json:"created"`
	Currency                  string      `json:"currency"`
	Description               string      `json:"description"`
	Destination               string      `json:"destination"`
	FailureBalanceTransaction interface{} `json:"failure_balance_transaction"`
	FailureCode               interface{} `json:"failure_code"`
	FailureMessage            interface{} `json:"failure_message"`
	Livemode                  bool        `json:"livemode"`
	Metadata                  struct {
	} `json:"metadata"`
	Method              string      `json:"method"`
	OriginalPayout      interface{} `json:"original_payout"`
	ReversedBy          interface{} `json:"reversed_by"`
	SourceType          string      `json:"source_type"`
	StatementDescriptor interface{} `json:"statement_descriptor"`
	Status              string      `json:"status"`
	Type                string      `json:"type"`
}
