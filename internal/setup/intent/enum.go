package intent

import "database/sql/driver"

type Status string

const (
	StatusRequiresPaymentMethod Status = "requires_payment_method"
	StatusRequiresConfirmation  Status = "requires_confirmation"
	StatusRequiresAction        Status = "requires_action"
	StatusProcessing            Status = "processing"
	StatusSucceeded             Status = "succeeded"
	StatusFailed                Status = "failed"
	StatusAbandoned             Status = "abandoned"
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

type CancellationReason string

const (
	CancellationReasonAbandoned           CancellationReason = "abandoned"
	CancellationReasonRequestedByCustomer CancellationReason = "requested_by_customer"
	CancellationReasonDuplicate           CancellationReason = "duplicate"
)

func (ct *CancellationReason) Scan(value interface{}) error {
	*ct = CancellationReason(value.(string))
	return nil
}

func (ct CancellationReason) Value() (driver.Value, error) {
	return string(ct), nil
}

type Type string

const (
	TypeRedirectToUrl           Type = "redirect_to_url"
	TypeUseStripeSdk            Type = "use_stripe_sdk"
	TypeAlipayHandleRedirect    Type = "alipay_handle_redirect"
	TypeOxxoDisplayDetails      Type = "oxxo_display_details"
	TypeVerifyWithMicrodeposits Type = "verify_with_microdeposits"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}

type MicrodepositType string

const (
	MicrodepositTypeDescriptorCode MicrodepositType = "descriptor_code"
	MicrodepositTypeAmounts        MicrodepositType = "amounts"
)

func (ct *MicrodepositType) Scan(value interface{}) error {
	*ct = MicrodepositType(value.(string))
	return nil
}

func (ct MicrodepositType) Value() (driver.Value, error) {
	return string(ct), nil
}
