package account

type AccountTosAcceptance struct {
	Date             int    `json:"date,omitempty"`
	Ip               string `json:"ip,omitempty"`
	ServiceAgreement string `json:"service_agreement,omitempty"`
	UserAgent        string `json:"user_agent,omitempty"`
}
