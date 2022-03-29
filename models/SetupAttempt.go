package models

type SetupAttempt struct {
	ID                   string      `json:"id"`
	Object               string      `json:"object"`
	Application          interface{} `json:"application"`
	Created              int         `json:"created"`
	Customer             interface{} `json:"customer"`
	Livemode             bool        `json:"livemode"`
	OnBehalfOf           interface{} `json:"on_behalf_of"`
	PaymentMethod        string      `json:"payment_method"`
	PaymentMethodDetails struct {
		Card struct {
			ThreeDSecure interface{} `json:"three_d_secure"`
		} `json:"card"`
		Type string `json:"type"`
	} `json:"payment_method_details"`
	SetupError  interface{} `json:"setup_error"`
	SetupIntent string      `json:"setup_intent"`
	Status      string      `json:"status"`
	Usage       string      `json:"usage"`
}