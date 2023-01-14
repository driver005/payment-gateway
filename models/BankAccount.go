package models

type BankAccount struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	AccountHolderName string      `json:"account_holder_name"`
	AccountHolderType string      `json:"account_holder_type"`
	AccountType       interface{} `json:"account_type"`
	BankName          string      `json:"bank_name"`
	Country           string      `json:"country"`
	Currency          string      `json:"currency"`
	Customer          interface{} `json:"customer"`
	Fingerprint       string      `json:"fingerprint"`
	Last4             string      `json:"last4"`
	Metadata          struct {
	} `json:"metadata"`
	RoutingNumber string `json:"routing_number"`
	Status        string `json:"status"`
}
