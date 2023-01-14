package models

import "database/sql/driver"

// The status of the Price List
type DiscountConditionType string

// Defines values for DiscountConditionType.
const (
	DiscountConditionTypeRroducts           DiscountConditionType = "products"
	DiscountConditionTypeProductTypes       DiscountConditionType = "product_types"
	DiscountConditionTypeProductCollections DiscountConditionType = "product_collections"
	DiscountConditionTypeProductTags        DiscountConditionType = "product_tags"
	DiscountConditionTypeCustomerGroups     DiscountConditionType = "customer_groups"
)

func (pl *DiscountConditionType) Scan(value interface{}) error {
	*pl = DiscountConditionType(value.([]byte))
	return nil
}

func (pl DiscountConditionType) Value() (driver.Value, error) {
	return string(pl), nil
}
