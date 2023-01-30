package methods

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&SetupAttemptPaymentMethodDetails{},
		&SetupAttemptPaymentMethodDetailsBancontact{},
		&SetupAttemptPaymentMethodDetailsCard{},
		&SetupAttemptPaymentMethodDetailsCardPresent{},
		&SetupAttemptPaymentMethodDetailsIdeal{},
		&SetupAttemptPaymentMethodDetailsSofort{},
	)
	if err != nil {
		panic(err)
	}
}
