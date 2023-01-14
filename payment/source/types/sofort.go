package types

// SourceTypeSofort struct for SourceTypeSofort
type SourceTypeSofort struct {
	BankCode string `json:"bank_code,omitempty"`
	BankName string `json:"bank_name,omitempty"`
	Bic string `json:"bic,omitempty"`
	Country string `json:"country,omitempty"`
	IbanLast4 string `json:"iban_last4,omitempty"`
	PreferredLanguage string `json:"preferred_language,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}
