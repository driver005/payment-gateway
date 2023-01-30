package fundingInstruction

import "github.com/driver005/gateway/core"

// FundingInstructionsBankTransferSortCodeRecord Sort Code Records contain U.K. bank account details per the sort code format.
type FundingInstructionsBankTransferSortCodeRecord struct {
	core.Model
	// The name of the person or business that owns the bank account
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The account number
	AccountNumber string `json:"account_number,omitempty"`
	// The six-digit sort code
	SortCode string `json:"sort_code,omitempty"`
}
