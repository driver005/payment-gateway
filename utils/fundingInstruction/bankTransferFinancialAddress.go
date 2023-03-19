package fundingInstruction

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// FundingInstructionsBankTransferFinancialAddress FinancialAddresses contain identifying information that resolves to a FinancialAccount.
type FundingInstructionsBankTransferFinancialAddress struct {
	core.Model

	Iban     *FundingInstructionsBankTransferIbanRecord     `json:"iban,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SortCode *FundingInstructionsBankTransferSortCodeRecord `json:"sort_code,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Spei     *FundingInstructionsBankTransferSpeiRecord     `json:"spei,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// The payment networks supported by this FinancialAddress
	SupportedNetworks pq.StringArray `json:"supported_networks,omitempty" database:"type:varchar(64)[]"`
	// The type of financial address
	Type   string                                       `json:"type,omitempty"`
	Zengin *FundingInstructionsBankTransferZenginRecord `json:"zengin,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
