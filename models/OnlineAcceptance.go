package models

// OnlineAcceptance
type OnlineAcceptance struct {
	// The IP address from which the Mandate was accepted by the customer.
	IpAddress string `json:"ip_address,omitempty"`
	// The user agent of the browser from which the Mandate was accepted by the customer.
	UserAgent string `json:"user_agent,omitempty"`
}
