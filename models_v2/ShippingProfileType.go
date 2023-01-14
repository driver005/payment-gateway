package models

import "database/sql/driver"

// The type of the Shipping Profile, may be `default`, `gift_card` or `custom`.
type ShippingProfileType string

// Defines values for ShippingProfileType.
const (
	ShippingProfileTypeCustom   ShippingProfileType = "custom"
	ShippingProfileTypeDefault  ShippingProfileType = "default"
	ShippingProfileTypeGiftCard ShippingProfileType = "gift_card"
)

func (sp *ShippingProfileType) Scan(value interface{}) error {
	*sp = ShippingProfileType(value.([]byte))
	return nil
}

func (sp ShippingProfileType) Value() (driver.Value, error) {
	return string(sp), nil
}
