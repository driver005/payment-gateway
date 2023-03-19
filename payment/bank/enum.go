package bank

import "database/sql/driver"

type AccountHolderType string

const (
	AccountHolderTypeIndividual AccountHolderType = "individual"
	AccountHolderTypeCompany    AccountHolderType = "company"
)

func (ct *AccountHolderType) Scan(value interface{}) error {
	*ct = AccountHolderType(value.(string))
	return nil
}

func (ct AccountHolderType) Value() (driver.Value, error) {
	return string(ct), nil
}

type AccountType string

const (
	AccountTypeChecking AccountType = "checking"
	AccountTypeSavings  AccountType = "savings"
	AccountTypeFutsu    AccountType = "futsu"
	AccountTypeToza     AccountType = "toza"
)

func (ct *AccountType) Scan(value interface{}) error {
	*ct = AccountType(value.(string))
	return nil
}

func (ct AccountType) Value() (driver.Value, error) {
	return string(ct), nil
}

type Status string

const (
	StatusNew                Status = "new"
	StatusValidated          Status = "validated"
	StatusVerified           Status = "verified"
	StatusVerificationFailed Status = "verification_failed"
	StatusErrored            Status = "errored"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}
