package models

type Tocken struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Card   struct {
		ID                string      `json:"id"`
		Object            string      `json:"object"`
		AddressCity       interface{} `json:"address_city"`
		AddressCountry    interface{} `json:"address_country"`
		AddressLine1      interface{} `json:"address_line1"`
		AddressLine1Check interface{} `json:"address_line1_check"`
		AddressLine2      interface{} `json:"address_line2"`
		AddressState      interface{} `json:"address_state"`
		AddressZip        interface{} `json:"address_zip"`
		AddressZipCheck   interface{} `json:"address_zip_check"`
		Brand             string      `json:"brand"`
		Country           string      `json:"country"`
		CvcCheck          string      `json:"cvc_check"`
		DynamicLast4      interface{} `json:"dynamic_last4"`
		ExpMonth          int         `json:"exp_month"`
		ExpYear           int         `json:"exp_year"`
		Fingerprint       string      `json:"fingerprint"`
		Funding           string      `json:"funding"`
		Last4             string      `json:"last4"`
		Metadata          struct {
		} `json:"metadata"`
		Name               interface{} `json:"name"`
		TokenizationMethod interface{} `json:"tokenization_method"`
	} `json:"card"`
	ClientIP interface{} `json:"client_ip"`
	Created  int         `json:"created"`
	Livemode bool        `json:"livemode"`
	Type     string      `json:"type"`
	Used     bool        `json:"used"`
}
