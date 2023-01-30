package types

import "github.com/driver005/gateway/core"

// SourceTypeGiropay struct for SourceTypeGiropay
type SourceTypeGiropay struct {
	core.Model

	BankCode            string `json:"bank_code,omitempty"`
	BankName            string `json:"bank_name,omitempty"`
	Bic                 string `json:"bic,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}
