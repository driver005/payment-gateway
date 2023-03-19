package intent

import "database/sql/driver"

type Status string

const (
	StatusRequiresPaymentMethod Status = "requires_payment_method"
	StatusRequiresConfirmation  Status = "requires_confirmation"
	StatusRequiresAction        Status = "requires_action"
	StatusProcessing            Status = "processing"
	StatusRequiresCapture       Status = "requires_capture"
	StatusCanceled              Status = "canceled"
	StatusSucceeded             Status = "succeeded"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type CancellationReason string

const (
	CancellationReasonDuplicate           CancellationReason = "duplicate"
	CancellationReasonFraudulent          CancellationReason = "fraudulent"
	CancellationReasonRequestedByCustomer CancellationReason = "requested_by_customer"
	CancellationReasonAbandoned           CancellationReason = "abandoned"
	CancellationReasonFailedInvoice       CancellationReason = "failed_invoice"
	CancellationReasonVoidInvoice         CancellationReason = "void_invoice"
	CancellationReasonAutomatic           CancellationReason = "automatic"
)

func (ct *CancellationReason) Scan(value interface{}) error {
	*ct = CancellationReason(value.(string))
	return nil
}

func (ct CancellationReason) Value() (driver.Value, error) {
	return string(ct), nil
}

type CaptureMethod string

const (
	CaptureMethodAutomatic CaptureMethod = "automatic"
	CaptureMethodManual    CaptureMethod = "manual"
)

func (ct *CaptureMethod) Scan(value interface{}) error {
	*ct = CaptureMethod(value.(string))
	return nil
}

func (ct CaptureMethod) Value() (driver.Value, error) {
	return string(ct), nil
}

type ConfirmationMethod string

const (
	ConfirmationMethodAutomatic ConfirmationMethod = "automatic"
	ConfirmationMethodManual    ConfirmationMethod = "manual"
)

func (ct *ConfirmationMethod) Scan(value interface{}) error {
	*ct = ConfirmationMethod(value.(string))
	return nil
}

func (ct ConfirmationMethod) Value() (driver.Value, error) {
	return string(ct), nil
}

type SetupFutureUsage string

const (
	SetupFutureUsageOffSession SetupFutureUsage = "off_session"
	SetupFutureUsageOnSession  SetupFutureUsage = "on_session"
)

func (ct *SetupFutureUsage) Scan(value interface{}) error {
	*ct = SetupFutureUsage(value.(string))
	return nil
}

func (ct SetupFutureUsage) Value() (driver.Value, error) {
	return string(ct), nil
}
