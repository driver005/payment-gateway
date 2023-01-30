package fundingInstruction

import "github.com/driver005/gateway/core"

// FundingInstructionsBankTransfer
type FundingInstructionsBankTransfer struct {
	core.Model
	// The country of the bank account to fund
	Country string `json:"country,omitempty"`
	// A list of financial addresses that can be used to fund a particular balance
	FinancialAddresses []FundingInstructionsBankTransferFinancialAddress `json:"financial_addresses,omitempty" database:"foreignKey:id"`
	// The bank_transfer type
	Type string `json:"type,omitempty"`
}
