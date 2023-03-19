package card

import "database/sql/driver"

type AddressZipCheck string

const (
	AddressZipCheckPass        AddressZipCheck = "pass"
	AddressZipCheckFail        AddressZipCheck = "fail"
	AddressZipCheckUnavailable AddressZipCheck = "unavailable"
	AddressZipCheckUnchecked   AddressZipCheck = "unchecked"
)

func (ct *AddressZipCheck) Scan(value interface{}) error {
	*ct = AddressZipCheck(value.(string))
	return nil
}

func (ct AddressZipCheck) Value() (driver.Value, error) {
	return string(ct), nil
}

type CvcCheck string

const (
	CvcCheckPass        CvcCheck = "pass"
	CvcCheckFail        CvcCheck = "fail"
	CvcCheckUnavailable CvcCheck = "unavailable"
	CvcCheckUnchecked   CvcCheck = "unchecked"
)

func (ct *CvcCheck) Scan(value interface{}) error {
	*ct = CvcCheck(value.(string))
	return nil
}

func (ct CvcCheck) Value() (driver.Value, error) {
	return string(ct), nil
}

type AddressLine1Check string

const (
	AddressLine1CheckPass        AddressLine1Check = "pass"
	AddressLine1CheckFail        AddressLine1Check = "fail"
	AddressLine1CheckUnavailable AddressLine1Check = "unavailable"
	AddressLine1CheckUnchecked   AddressLine1Check = "unchecked"
)

func (ct *AddressLine1Check) Scan(value interface{}) error {
	*ct = AddressLine1Check(value.(string))
	return nil
}

func (ct AddressLine1Check) Value() (driver.Value, error) {
	return string(ct), nil
}

type Brand string

const (
	BrandAmericanExpress Brand = "American Express"
	BrandDinersClub      Brand = "Diners Club"
	BrandDiscover        Brand = "Discover"
	BrandJCB             Brand = "JCB"
	BrandMasterCard      Brand = "MasterCard"
	BrandUnionPay        Brand = "UnionPay"
	BrandVisa            Brand = "Visa"
	BrandUnknown         Brand = "Unknown"
)

func (ct *Brand) Scan(value interface{}) error {
	*ct = Brand(value.(string))
	return nil
}

func (ct Brand) Value() (driver.Value, error) {
	return string(ct), nil
}

type Funding string

const (
	FundingCredit  Funding = "credit"
	FundingDebit   Funding = "debit"
	FundingPrepaid Funding = "prepaid"
	FundingUnknown Funding = "unknown"
)

func (ct *Funding) Scan(value interface{}) error {
	*ct = Funding(value.(string))
	return nil
}

func (ct Funding) Value() (driver.Value, error) {
	return string(ct), nil
}

type TokenizationMethod string

const (
	TokenizationMethodAndroidPay   TokenizationMethod = "android_pay"
	TokenizationMethodApplePay     TokenizationMethod = "apple_pay"
	TokenizationMethodMasterpass   TokenizationMethod = "masterpass"
	TokenizationMethodVisaCheckout TokenizationMethod = "visa_checkout"
	TokenizationMethodUnknown      TokenizationMethod = "unknown"
)

func (ct *TokenizationMethod) Scan(value interface{}) error {
	*ct = TokenizationMethod(value.(string))
	return nil
}

func (ct TokenizationMethod) Value() (driver.Value, error) {
	return string(ct), nil
}
