package fundingInstruction

import "github.com/driver005/gateway/core"

// FundingInstructionsBankTransferSpeiRecord SPEI Records contain Mexico bank account details per the SPEI format.
type FundingInstructionsBankTransferSpeiRecord struct {
	core.Model
	// The three-digit bank code
	BankCode string `json:"bank_code,omitempty"`
	// The short banking institution name
	BankName string `json:"bank_name,omitempty"`
	// The CLABE number
	Clabe string `json:"clabe,omitempty"`
}
