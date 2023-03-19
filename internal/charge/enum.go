package charge

import "database/sql/driver"

type Type string

const (
	TypeAuthorized     Type = "authorized"
	TypeManualReview   Type = "manual_review"
	TypeIssuerDeclined Type = "issuer_declined"
	TypeBlocked        Type = "blocked"
	TypeInvalid        Type = "invalid"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value, error) {
	return string(ct), nil
}

type Status string

const (
	StatusSucceeded Status = "succeeded"
	StatusPending   Status = "pending"
	StatusFailed    Status = "failed"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type NetworkStatus string

const (
	NetworkStatusApprovedByNetwork     NetworkStatus = "approved_by_network"
	NetworkStatusDeclinedByNetwork     NetworkStatus = "declined_by_network"
	NetworkStatusNotSentToNetwork      NetworkStatus = "not_sent_to_network"
	NetworkStatusReversedAfterApproval NetworkStatus = "reversed_after_approval"
)

func (ct *NetworkStatus) Scan(value interface{}) error {
	*ct = NetworkStatus(value.(string))
	return nil
}

func (ct NetworkStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Reason string

const (
	ReasonHighestRiskLevel  Reason = "highest_risk_level"
	ReasonElevatedRiskLevel Reason = "elevated_risk_level"
)

func (ct *Reason) Scan(value interface{}) error {
	*ct = Reason(value.(string))
	return nil
}

func (ct Reason) Value() (driver.Value, error) {
	return string(ct), nil
}

type RiskLevel string

const (
	RiskLevelNormal      RiskLevel = "normal"
	RiskLevelElevated    RiskLevel = "elevated"
	RiskLevelHighest     RiskLevel = "highest"
	RiskLevelNotAssessed RiskLevel = "not_assessed"
	RiskLevelUnknown     RiskLevel = "unknown"
)

func (ct *RiskLevel) Scan(value interface{}) error {
	*ct = RiskLevel(value.(string))
	return nil
}

func (ct RiskLevel) Value() (driver.Value, error) {
	return string(ct), nil
}
