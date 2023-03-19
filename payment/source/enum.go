package source

import "database/sql/driver"

type Status string

const (
	StatusCanceled   Status = "canceled"
	StatusChargeable Status = "chargeable"
	StatusConsumed   Status = "consumed"
	StatusFailed     Status = "failed"
	StatusPending    Status = "pending"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type Type string

const (
	TypeAchreusableTransfer Type = "ach_reusable_transfer"
	TypeAchDebit            Type = "ach_debit"
	TypeAlipay              Type = "alipay"
	TypeBancontact          Type = "bancontact"
	TypeCard                Type = "card"
	TypeCardPresent         Type = "card_present"
	TypeEps                 Type = "eps"
	TypeGiropay             Type = "giropay"
	TypeIdeal               Type = "ideal"
	TypeMultibanco          Type = "multibanco"
	TypeKlarna              Type = "klarna"
	TypeP24                 Type = "p24"
	TypeSepaDebit           Type = "sepa_debit"
	TypeSofort              Type = "sofort"
	TypeThreeDSecure        Type = "three_d_secure"
	TypeWechat              Type = "wechat"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}

type Flow string

const (
	FlowRedirect         Flow = "redirect"
	FlowReceiver         Flow = "receiver"
	FlowCodeVerification Flow = "code_verification"
	FlowNone             Flow = "none"
)

func (ct *Flow) Scan(value interface{}) error {
	*ct = Flow(value.(string))
	return nil
}

func (ct Flow) Value() (driver.Value, error) {
	return string(ct), nil
}

type RefundAttributesMethod string

const (
	RefundAttributesMethodEmail  RefundAttributesMethod = "email"
	RefundAttributesMethodManual RefundAttributesMethod = "manual"
	RefundAttributesMethodNone   RefundAttributesMethod = "none"
)

func (ct *RefundAttributesMethod) Scan(value interface{}) error {
	*ct = RefundAttributesMethod(value.(string))
	return nil
}

func (ct RefundAttributesMethod) Value() (driver.Value, error) {
	return string(ct), nil
}

type RefundAttributesStatus string

const (
	RefundAttributesStatusMissing   RefundAttributesStatus = "missing"
	RefundAttributesStatusRequested RefundAttributesStatus = "requested"
	RefundAttributesStatusAvailable RefundAttributesStatus = "available"
)

func (ct *RefundAttributesStatus) Scan(value interface{}) error {
	*ct = RefundAttributesStatus(value.(string))
	return nil
}

func (ct RefundAttributesStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Usage string

const (
	UsageReusable  Usage = "reusable"
	UsageSingleUse Usage = "single_use"
)

func (ct *Usage) Scan(value interface{}) error {
	*ct = Usage(value.(string))
	return nil
}

func (ct Usage) Value() (driver.Value, error) {
	return string(ct), nil
}

type ItemType string

const (
	ItemTypeSku      ItemType = "sku"
	ItemTypeTax      ItemType = "tax"
	ItemTypeShipping ItemType = "shipping"
)

func (ct *ItemType) Scan(value interface{}) error {
	*ct = ItemType(value.(string))
	return nil
}

func (ct ItemType) Value() (driver.Value, error) {
	return string(ct), nil
}

type CodeStatus string

const (
	CodeStatusPending           CodeStatus = "pending"
	CodeStatusAttemptsRemaining CodeStatus = "attempts_remaining"
	CodeStatusSucceeded         CodeStatus = "succeeded"
	CodeStatusFailed            CodeStatus = "failed "
)

func (ct *CodeStatus) Scan(value interface{}) error {
	*ct = CodeStatus(value.(string))
	return nil
}

func (ct CodeStatus) Value() (driver.Value, error) {
	return string(ct), nil
}
