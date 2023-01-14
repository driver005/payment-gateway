package fundingInstruction

// FundingInstructionsBankTransferFinancialAddress FinancialAddresses contain identifying information that resolves to a FinancialAccount.
type FundingInstructionsBankTransferFinancialAddress struct {
	Iban     *FundingInstructionsBankTransferIbanRecord     `json:"iban,omitempty"`
	SortCode *FundingInstructionsBankTransferSortCodeRecord `json:"sort_code,omitempty"`
	Spei     *FundingInstructionsBankTransferSpeiRecord     `json:"spei,omitempty"`
	// The payment networks supported by this FinancialAddress
	SupportedNetworks []string `json:"supported_networks,omitempty"`
	// The type of financial address
	Type   string                                       `json:"type"`
	Zengin *FundingInstructionsBankTransferZenginRecord `json:"zengin,omitempty"`
}
