package types

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&SourceTypeAchCreditTransfer{},
		&SourceTypeAchDebit{},
		&SourceTypeAcssDebit{},
		&SourceTypeAlipay{},
		&SourceTypeAuBecsDebit{},
		&SourceTypeBancontact{},
		&SourceTypeCard{},
		&SourceTypeCardPresent{},
		&SourceTypeGiropay{},
		&SourceTypeIdeal{},
		&SourceTypeKlarna{},
		&SourceTypeMultibanco{},
		&SourceTypeP24{},
		&SourceTypeSepaDebit{},
		&SourceTypeSofort{},
		&SourceTypeThreeDSecure{},
		&SourceTypeWechat{},
	)
	if err != nil {
		panic(err)
	}
}
