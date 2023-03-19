package credit

import "database/sql/driver"

type Reason string

const (
	ReasonDuplicate   Reason = "duplicate"
	ReasonFraudulent  Reason = "fraudulent"
	ReasonOrderChange Reason = "order_change"
)

func (ct *Reason) Scan(value interface{}) error {
	*ct = Reason(value.(string))
	return nil
}

func (ct Reason) Value() (driver.Value, error) {
	return string(ct), nil
}

type Status string

const (
	StatusIssued Status = "issued"
	StatusVoid   Status = "void"
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
	TypePrePayment  Type = "pre_payment"
	TypePostPayment Type = "post_payment"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}

type LineItemType string

const (
	LineItemTypeInvoiceLineItemType LineItemType = "invoice_line_item"
	LineItemTypeCustomLineItemType  LineItemType = "custom_line_item"
)

func (ct *LineItemType) Scan(value interface{}) error {
	*ct = LineItemType(value.(string))
	return nil
}

func (ct LineItemType) Value() (driver.Value, error) {
	return string(ct), nil
}
