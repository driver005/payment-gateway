package models

type Balance struct {
	Object    string `json:"object"`
	Available []struct {
		Amount      int    `json:"amount"`
		Currency    string `json:"currency"`
		SourceTypes struct {
			BankAccount int `json:"bank_account"`
			Card        int `json:"card"`
		} `json:"source_types"`
	} `json:"available"`
	ConnectReserved []struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"connect_reserved"`
	Livemode bool `json:"livemode"`
	Pending  []struct {
		Amount      int    `json:"amount"`
		Currency    string `json:"currency"`
		SourceTypes struct {
			BankAccount int `json:"bank_account"`
			Card        int `json:"card"`
		} `json:"source_types"`
	} `json:"pending"`
}
