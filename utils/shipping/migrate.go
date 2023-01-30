package shipping

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&ShippingRate{},
		&ShippingRateCurrencyOption{},
		&ShippingRateDeliveryEstimate{},
		&ShippingRateDeliveryEstimateBound{},
		&ShippingRateFixedAmount{},
	)
	if err != nil {
		panic(err)
	}
}
