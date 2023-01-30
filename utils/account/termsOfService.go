package account

type AccountTermsOfService struct {
	Date      int    `json:"date,omitempty"`
	Ip        string `json:"ip,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
}
