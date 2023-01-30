package types

import "github.com/driver005/gateway/core"

// SourceTypeBancontact struct for SourceTypeBancontact
type SourceTypeBancontact struct {
	core.Model

	BankCode            string `json:"bank_code,omitempty"`
	BankName            string `json:"bank_name,omitempty"`
	Bic                 string `json:"bic,omitempty"`
	IbanLast4           string `json:"iban_last4,omitempty"`
	PreferredLanguage   string `json:"preferred_language,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}
