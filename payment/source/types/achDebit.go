package types

// SourceTypeAchDebit struct for SourceTypeAchDebit
type SourceTypeAchDebit struct {
	BankName string `json:"bank_name,omitempty"`
	Country string `json:"country,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Last4 string `json:"last4,omitempty"`
	RoutingNumber string `json:"routing_number,omitempty"`
	Type string `json:"type,omitempty"`
}
