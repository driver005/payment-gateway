package models

type CustomerPortal struct {
	ID            string      `json:"id"`
	Object        string      `json:"object"`
	Configuration string      `json:"configuration"`
	Created       int         `json:"created"`
	Customer      string      `json:"customer"`
	Livemode      bool        `json:"livemode"`
	Locale        interface{} `json:"locale"`
	OnBehalfOf    interface{} `json:"on_behalf_of"`
	ReturnURL     string      `json:"return_url"`
	URL           string      `json:"url"`
}