package fundingInstruction

import "github.com/driver005/gateway/core"

// FundingInstructionsBankTransferFinancialAddress FinancialAddresses contain identifying information that resolves to a FinancialAccount.
type FundingInstructionsBankTransferFinancialAddress struct {
	core.Model

	Iban     *FundingInstructionsBankTransferIbanRecord     `json:"iban,omitempty" database:"foreignKey:id"`
	SortCode *FundingInstructionsBankTransferSortCodeRecord `json:"sort_code,omitempty" database:"foreignKey:id"`
	Spei     *FundingInstructionsBankTransferSpeiRecord     `json:"spei,omitempty" database:"foreignKey:id"`
	// The payment networks supported by this FinancialAddress
	SupportedNetworks []string `json:"supported_networks,omitempty" database:"type:text[]"`
	// The type of financial address
	Type   string                                       `json:"type,omitempty"`
	Zengin *FundingInstructionsBankTransferZenginRecord `json:"zengin,omitempty" database:"foreignKey:id"`
}
