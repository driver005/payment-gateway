package attempt

import "database/sql/driver"

type Status string

const (
	StatusRequiresConfirmation Status = "requires_confirmation"
	StatusRequiresAction       Status = "requires_action"
	StatusProcessing           Status = "processing"
	StatusSucceeded            Status = "succeeded"
	StatusFailed               Status = "failed"
	StatusAbandoned            Status = "abandoned"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type Usage string

const (
	UsageOffSession Usage = "off_session"
	UsageOnSession  Usage = "on_session"
)

func (ct *Usage) Scan(value interface{}) error {
	*ct = Usage(value.(string))
	return nil
}

func (ct Usage) Value() (driver.Value, error) {
	return string(ct), nil
}
