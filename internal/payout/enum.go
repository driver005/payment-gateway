package payout

import "database/sql/driver"

type Status string

const (
	StatusPaid      Status = "paid"
	StatusPending   Status = "pending"
	StatusInTransit Status = "in_transit"
	StatusCanceled  Status = "canceled"
	StatusFailed    Status = "failed"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type Method string

const (
	MethodStandard Method = "standard"
	MethodInstant  Method = "instant"
)

func (ct *Method) Scan(value interface{}) error {
	*ct = Method(value.(string))
	return nil
}

func (ct Method) Value() (driver.Value, error) {
	return string(ct), nil
}

type SourceType string

const (
	SourceTypeCard        SourceType = "card"
	SourceTypeFpx         SourceType = "fpx"
	SourceTypeBankAccount SourceType = "bank_account"
)

func (ct *SourceType) Scan(value interface{}) error {
	*ct = SourceType(value.(string))
	return nil
}

func (ct SourceType) Value() (driver.Value, error) {
	return string(ct), nil
}

type Type string

const (
	TypeBankAccount Type = "bank_account"
	TypeCard        Type = "card"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}
