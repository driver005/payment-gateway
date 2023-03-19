package price

import "database/sql/driver"

type AggregateUsage string

const (
	AggregateUsageSum              AggregateUsage = "sum"
	AggregateUsageLastDuringPeriod AggregateUsage = "last_during_period"
	AggregateUsageLastEver         AggregateUsage = "last_ever"
	AggregateUsageMax              AggregateUsage = "max"
)

func (ct *AggregateUsage) Scan(value interface{}) error {
	*ct = AggregateUsage(value.(string))
	return nil
}

func (ct AggregateUsage) Value() (driver.Value, error) {
	return string(ct), nil
}

type Type string

const (
	TypeOneTime   Type = "one_time"
	TypeRecurring Type = "recurring"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}

type Interval string

const (
	IntervalMonth Interval = "month"
	IntervalYear  Interval = "year"
	IntervalWeek  Interval = "week"
	IntervalDay   Interval = "day"
)

func (ct *Interval) Scan(value interface{}) error {
	*ct = Interval(value.(string))
	return nil
}

func (ct Interval) Value() (driver.Value, error) {
	return string(ct), nil
}

type UsageType string

const (
	UsageTypeMetered  UsageType = "metered"
	UsageTypeLicensed UsageType = "licensed"
)

func (ct *UsageType) Scan(value interface{}) error {
	*ct = UsageType(value.(string))
	return nil
}

func (ct UsageType) Value() (driver.Value, error) {
	return string(ct), nil
}

type BillingScheme string

const (
	BillingSchemePerUnit BillingScheme = "per_unit"
	BillingSchemeTiered  BillingScheme = "tiered"
)

func (ct *BillingScheme) Scan(value interface{}) error {
	*ct = BillingScheme(value.(string))
	return nil
}

func (ct BillingScheme) Value() (driver.Value, error) {
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

type TiersMode string

const (
	TiersModeGraduated TiersMode = "graduated"
	TiersModeVolume    TiersMode = "volume"
)

func (ct *TiersMode) Scan(value interface{}) error {
	*ct = TiersMode(value.(string))
	return nil
}

func (ct TiersMode) Value() (driver.Value, error) {
	return string(ct), nil
}

type Round string

const (
	RoundUp   Round = "up"
	RoundDown Round = "down"
)

func (ct *Round) Scan(value interface{}) error {
	*ct = Round(value.(string))
	return nil
}

func (ct Round) Value() (driver.Value, error) {
	return string(ct), nil
}
