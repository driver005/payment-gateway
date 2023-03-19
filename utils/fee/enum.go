package fee

import "database/sql/driver"

type Type string

const (
	TypeApplicationFee Type = "application_fee"
	TypeStripeFee      Type = "stripe_fee"
	TypeTax            Type = "tax"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}
