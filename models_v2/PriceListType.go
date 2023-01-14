package models

import "database/sql/driver"

// The type of Price List. This can be one of either `sale` or `override`.
type PriceListType string

const (
	Override PriceListType = "override"
	Sale     PriceListType = "sale"
)

func (pl *PriceListType) Scan(value interface{}) error {
	*pl = PriceListType(value.([]byte))
	return nil
}

func (pl PriceListType) Value() (driver.Value, error) {
	return string(pl), nil
}
