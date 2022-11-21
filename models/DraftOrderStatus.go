package models

import "database/sql/driver"

// The status of the Price List
type DraftOrderStatus string

// Defines values for DraftOrderStatus.
const (
	DraftOrderStatusCompleted DraftOrderStatus = "completed"
	DraftOrderStatusOpen      DraftOrderStatus = "open"
)

func (pl *DraftOrderStatus) Scan(value interface{}) error {
	*pl = DraftOrderStatus(value.([]byte))
	return nil
}

func (pl DraftOrderStatus) Value() (driver.Value, error) {
	return string(pl), nil
}
