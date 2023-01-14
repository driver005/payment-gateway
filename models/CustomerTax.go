package models

type CustomerTax struct {
	ID           string `json:"id"`
	Object       string `json:"object"`
	Country      string `json:"country"`
	Created      int    `json:"created"`
	Customer     string `json:"customer"`
	Livemode     bool   `json:"livemode"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	Verification struct {
		Status          string      `json:"status"`
		VerifiedAddress interface{} `json:"verified_address"`
		VerifiedName    interface{} `json:"verified_name"`
	} `json:"verification"`
}
