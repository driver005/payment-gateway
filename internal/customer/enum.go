package customer

import "database/sql/driver"

type TaxExempt string

const (
	TaxExemptNone    TaxExempt = "none"
	TaxExemptExempt  TaxExempt = "exempt"
	TaxExemptReverse TaxExempt = "reverse"
)

func (ct *TaxExempt) Scan(value interface{}) error {
	*ct = TaxExempt(value.(string))
	return nil
}

func (ct TaxExempt) Value() (driver.Value, error) {
	return string(ct), nil
}
