package fundingInstruction

// FundingInstructionsBankTransfer
type FundingInstructionsBankTransfer struct {
	// The country of the bank account to fund
	Country string `json:"country"`
	// A list of financial addresses that can be used to fund a particular balance
	FinancialAddresses []FundingInstructionsBankTransferFinancialAddress `json:"financial_addresses"`
	// The bank_transfer type
	Type string `json:"type"`
}
