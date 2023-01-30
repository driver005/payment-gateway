package account

type CardIssuingAccountTermsOfService struct {
	// The Unix timestamp marking when the account representative accepted the service agreement.
	Date int `json:"date,omitempty"`
	// The IP address from which the account representative accepted the service agreement.
	Ip string `json:"ip,omitempty"`
	// The user agent of the browser from which the account representative accepted the service agreement.
	UserAgent string `json:"user_agent,omitempty"`
}

// AccountCardIssuingSettings
type AccountCardIssuingSettings struct {
	TosAcceptance *CardIssuingAccountTermsOfService `json:"tos_acceptance,omitempty"`
}
