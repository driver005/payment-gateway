package coupon

import "database/sql/driver"

type Duration string

const (
	DurationOnce      Duration = "once"
	DurationRepeating Duration = "repeating"
	DurationForever   Duration = "forever"
)

func (ct *Duration) Scan(value interface{}) error {
	*ct = Duration(value.(string))
	return nil
}

func (ct Duration) Value() (driver.Value, error) {
	return string(ct), nil
}
