package methods

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&MandatePaymentMethodDetails{},
		&MandateAcssDebit{},
		&MandateAuBecsDebit{},
		&MandateBacsDebit{},
		&MandateBlik{},
		&MandateCard{},
		&MandateLink{},
		&MandateOptionsOffSessionDetailsBlik{},
		&MandateOptionsParam{},
		&MandateSepaDebit{},
		&MandateSingleUse{},
		&MandateUsBankAccount{},
	)
	if err != nil {
		panic(err)
	}
}
