package models

import "database/sql/driver"

// The status of the Price List
type BatchJobStatus string

// Defines values for BatchJobStatus.
const (
	BatchJobStatusCreated      BatchJobStatus = "created"
	BatchJobStatusPreProcessed BatchJobStatus = "pre_processed"
	BatchJobStatusConfirmed    BatchJobStatus = "confirmed"
	BatchJobStatusProcessing   BatchJobStatus = "processing"
	BatchJobStatusCompleted    BatchJobStatus = "completed"
	BatchJobStatusCanceled     BatchJobStatus = "canceled"
	BatchJobStatusFailed       BatchJobStatus = "failed"
)

func (pl *BatchJobStatus) Scan(value interface{}) error {
	*pl = BatchJobStatus(value.([]byte))
	return nil
}

func (pl BatchJobStatus) Value() (driver.Value, error) {
	return string(pl), nil
}
