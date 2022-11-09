package models

import "database/sql/driver"

// The status of the Price List
type PriceListStatus string

// Defines values for PriceListStatus.
const (
	PriceListStatusActive PriceListStatus = "active"
	PriceListStatusDraft  PriceListStatus = "draft"
)

func (pl *PriceListStatus) Scan(value interface{}) error {
	*pl = PriceListStatus(value.([]byte))
	return nil
}

func (pl PriceListStatus) Value() (driver.Value, error) {
	return string(pl), nil
}
