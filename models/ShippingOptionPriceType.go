package models

import "database/sql/driver"

// The type of pricing calculation that is used when creatin Shipping Methods from the Shipping Option. Can be `flat_rate` for fixed prices or `calculated` if the Fulfillment Provider can provide price calulations.
type ShippingOptionPriceType string

// Defines values for ShippingOptionPriceType.
const (
	Calculated ShippingOptionPriceType = "calculated"
	FlatRate   ShippingOptionPriceType = "flat_rate"
)

func (so *ShippingOptionPriceType) Scan(value interface{}) error {
	*so = ShippingOptionPriceType(value.([]byte))
	return nil
}

func (so ShippingOptionPriceType) Value() (driver.Value, error) {
	return string(so), nil
}
