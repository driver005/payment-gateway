package models

type BalanceTransaction struct {
	ID           string      `json:"id"`
	Object       string      `json:"object"`
	Amount       int         `json:"amount"`
	AvailableOn  int         `json:"available_on"`
	Created      int         `json:"created"`
	Currency     string      `json:"currency"`
	Description  string      `json:"description"`
	ExchangeRate interface{} `json:"exchange_rate"`
	Fee          int         `json:"fee"`
	FeeDetails   []struct {
		Amount      int         `json:"amount"`
		Application interface{} `json:"application"`
		Currency    string      `json:"currency"`
		Description string      `json:"description"`
		Type        string      `json:"type"`
	} `json:"fee_details"`
	Net               int    `json:"net"`
	ReportingCategory string `json:"reporting_category"`
	Source            string `json:"source"`
	Status            string `json:"status"`
	Type              string `json:"type"`
}