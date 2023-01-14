package types

// SourceTypeAcssDebit struct for SourceTypeAcssDebit
type SourceTypeAcssDebit struct {
	BankAddressCity string `json:"bank_address_city,omitempty"`
	BankAddressLine1 string `json:"bank_address_line_1,omitempty"`
	BankAddressLine2 string `json:"bank_address_line_2,omitempty"`
	BankAddressPostalCode string `json:"bank_address_postal_code,omitempty"`
	BankName string `json:"bank_name,omitempty"`
	Category string `json:"category,omitempty"`
	Country string `json:"country,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Last4 string `json:"last4,omitempty"`
	RoutingNumber string `json:"routing_number,omitempty"`
}
