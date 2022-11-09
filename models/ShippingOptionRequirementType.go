package models

import "database/sql/driver"

// The type of the requirement, this defines how the value will be compared to the Cart's total. `min_subtotal` requirements define the minimum subtotal that is needed for the Shipping Option to be available, while the `max_subtotal` defines the maximum subtotal that the Cart can have for the Shipping Option to be available.
type ShippingOptionRequirementType string

// Defines values for ShippingOptionRequirementType.
const (
	MaxSubtotal ShippingOptionRequirementType = "max_subtotal"
	MinSubtotal ShippingOptionRequirementType = "min_subtotal"
)

func (so *ShippingOptionRequirementType) Scan(value interface{}) error {
	*so = ShippingOptionRequirementType(value.([]byte))
	return nil
}

func (so ShippingOptionRequirementType) Value() (driver.Value, error) {
	return string(so), nil
}
