package types

import "github.com/driver005/gateway/core"

// SourceTypeSepaDebit struct for SourceTypeSepaDebit
type SourceTypeSepaDebit struct {
	core.Model

	BankCode         string `json:"bank_code,omitempty"`
	BranchCode       string `json:"branch_code,omitempty"`
	Country          string `json:"country,omitempty"`
	Fingerprint      string `json:"fingerprint,omitempty"`
	Last4            string `json:"last4,omitempty"`
	MandateReference string `json:"mandate_reference,omitempty"`
	MandateUrl       string `json:"mandate_url,omitempty"`
}
