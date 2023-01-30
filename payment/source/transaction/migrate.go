package transaction

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&SourceTransactionAchCreditTransferData{},
		&SourceTransactionChfCreditTransferData{},
		&SourceTransactionGbpCreditTransferData{},
		&SourceTransactionPaperCheckData{},
		&SourceTransactionSepaCreditTransferData{},
		&SourceTransaction{},
	)
	if err != nil {
		panic(err)
	}
}
