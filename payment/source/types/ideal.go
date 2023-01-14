package types

// SourceTypeIdeal struct for SourceTypeIdeal
type SourceTypeIdeal struct {
	Bank string `json:"bank,omitempty"`
	Bic string `json:"bic,omitempty"`
	IbanLast4 string `json:"iban_last4,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}
