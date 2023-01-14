package fundingInstruction

// FundingInstructionsBankTransferSpeiRecord SPEI Records contain Mexico bank account details per the SPEI format.
type FundingInstructionsBankTransferSpeiRecord struct {
	// The three-digit bank code
	BankCode string `json:"bank_code"`
	// The short banking institution name
	BankName string `json:"bank_name"`
	// The CLABE number
	Clabe string `json:"clabe"`
}
