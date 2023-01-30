package fundingInstruction

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&FundingInstructions{},
		&FundingInstructionsBankTransfer{},
		&FundingInstructionsBankTransferFinancialAddress{},
		&FundingInstructionsBankTransferIbanRecord{},
		&FundingInstructionsBankTransferSortCodeRecord{},
		&FundingInstructionsBankTransferSpeiRecord{},
		&FundingInstructionsBankTransferZenginRecord{},
	)
	if err != nil {
		panic(err)
	}
}
