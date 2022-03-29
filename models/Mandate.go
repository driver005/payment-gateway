package models

type Mandate struct {
	ID                 string `json:"id"`
	Object             string `json:"object"`
	CustomerAcceptance struct {
		AcceptedAt int `json:"accepted_at"`
		Online     struct {
			IPAddress string `json:"ip_address"`
			UserAgent string `json:"user_agent"`
		} `json:"online"`
		Type string `json:"type"`
	} `json:"customer_acceptance"`
	Livemode bool `json:"livemode"`
	MultiUse struct {
	} `json:"multi_use"`
	PaymentMethod        string `json:"payment_method"`
	PaymentMethodDetails struct {
		SepaDebit struct {
			Reference string `json:"reference"`
			URL       string `json:"url"`
		} `json:"sepa_debit"`
		Type string `json:"type"`
	} `json:"payment_method_details"`
	Status string `json:"status"`
	Type   string `json:"type"`
}