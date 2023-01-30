package types

import "github.com/driver005/gateway/core"

// SourceTypeIdeal struct for SourceTypeIdeal
type SourceTypeIdeal struct {
	core.Model

	Bank                string `json:"bank,omitempty"`
	Bic                 string `json:"bic,omitempty"`
	IbanLast4           string `json:"iban_last4,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}
