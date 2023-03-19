package refund

import "database/sql/driver"

type Status string

const (
	StatusPending        Status = "pending"
	StatusRequiresAction Status = "requires_action"
	StatusSucceeded      Status = "succeeded"
	StatusFailed         Status = "failed"
	StatusCanceled       Status = "canceled"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type Reason string

const (
	ReasonDuplicate               Reason = "duplicate"
	ReasonFraudulent              Reason = "fraudulent"
	ReasonRequestedByCustomer     Reason = "requested_by_customer"
	ReasonExpiredUncapturedCharge Reason = "expired_uncaptured_charge"
)

func (ct *Reason) Scan(value interface{}) error {
	*ct = Reason(value.(string))
	return nil
}

func (ct Reason) Value() (driver.Value, error) {
	return string(ct), nil
}

type FailureReason string

const (
	FailureReasonLostOrStolenCard      FailureReason = "lost_or_stolen_card"
	FailureReasonExpiredOrCanceledCard FailureReason = "expired_or_canceled_card"
	FailureReasonUnknown               FailureReason = "unknown"
)

func (ct *FailureReason) Scan(value interface{}) error {
	*ct = FailureReason(value.(string))
	return nil
}

func (ct FailureReason) Value() (driver.Value, error) {
	return string(ct), nil
}
