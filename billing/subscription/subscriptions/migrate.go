package subscriptions

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&SubscriptionsResourcePaymentSettings{},
		&SubscriptionsResourcePauseCollection{},
		&SubscriptionsResourcePendingUpdate{},
	)
	if err != nil {
		panic(err)
	}
}
