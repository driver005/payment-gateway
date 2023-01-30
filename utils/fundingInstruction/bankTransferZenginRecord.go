package fundingInstruction

import "github.com/driver005/gateway/core"

// FundingInstructionsBankTransferZenginRecord Zengin Records contain Japan bank account details per the Zengin format.
type FundingInstructionsBankTransferZenginRecord struct {
	core.Model
	// The account holder name
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The account number
	AccountNumber string `json:"account_number,omitempty"`
	// The bank account type. In Japan, this can only be `futsu` or `toza`.
	AccountType string `json:"account_type,omitempty"`
	// The bank code of the account
	BankCode string `json:"bank_code,omitempty"`
	// The bank name of the account
	BankName string `json:"bank_name,omitempty"`
	// The branch code of the account
	BranchCode string `json:"branch_code,omitempty"`
	// The branch name of the account
	BranchName string `json:"branch_name,omitempty"`
}
