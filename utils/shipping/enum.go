package shipping

import "database/sql/driver"

type Type string

const (
	TypeFixedAmount Type = "fixed_amount"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}

type TaxBehavior string

const (
	TaxBehaviorInclusive   TaxBehavior = "inclusive"
	TaxBehaviorExclusive   TaxBehavior = "exclusive"
	TaxBehaviorUnspecified TaxBehavior = "unspecified"
)

func (ct *TaxBehavior) Scan(value interface{}) error {
	*ct = TaxBehavior(value.(string))
	return nil
}

func (ct TaxBehavior) Value() (driver.Value, error) {
	return string(ct), nil
}
