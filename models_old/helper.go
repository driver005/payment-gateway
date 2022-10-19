package models

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/driver005/database"
	"github.com/driver005/database/types"
	"github.com/driver005/gateway/helper"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Address struct {
	City       string `json:"city" db:"city"`
	State      string `json:"state" db:"states"`
	Country    string `json:"country" db:"country"`
	Line1      string `json:"line1" db:"line1"`
	Line1Check string `json:"line1_check" db:"line1_check"`
	Line2      string `json:"line2" db:"line2"`
	Zip        string `json:"zip" db:"zip"`
	ZipCheck   string `json:"zip_check" db:"zip_check"`
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.City, validation.Length(1, 50)),
		validation.Field(&a.Country, validation.Length(5, 50)),
		validation.Field(&a.Line1, validation.Length(1, 50)),
		validation.Field(&a.Line1Check, validation.In("pass", "fail", "unavailable", "unchecked")),
		validation.Field(&a.Line2, validation.Length(1, 50)),
		validation.Field(&a.State, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&a.Zip, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		validation.Field(&a.ZipCheck, validation.In("pass", "fail", "unavailable", "unchecked")),
	)
}

func (a Address) Interface() interface{} {
	return Address(a)
}

func (Address) DBDataType() string {
	return "addressType"
}

func (a Address) DBValue(ctx context.Context, db *database.DB) types.Expr {
	return types.Expr{
		SQL:  "ROW(?, ?, ?, ?, ?, ?, ?, ?)::addressType",
		Vars: []interface{}{a.City, a.Country, a.Line1, a.Line1Check, a.Line2, a.State, a.Zip, a.ZipCheck},
	}
}

func (a *Address) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	s := strings.Split(strings.Trim(string(bytes), "()"), ",")
	for i := 0; i < reflect.TypeOf(Address{}).NumField(); i++ {
		helper.SetField(&a, reflect.TypeOf(Address{}).Field(i).Name, s[i])
	}

	return nil
}

type Checks struct {
	AddressLine1Check      string `json:"address_line1_check"`
	AddressPostalCodeCheck string `json:"address_postal_code_check"`
	CvcCheck               string `json:"cvc_check"`
}

//Banks
type ThreeDSecure struct {
	AuthenticationFlow string `json:"authentication_flow"`
	Result             string `json:"result"`
	ResultReason       string `json:"result_reason"`
	Version            string `json:"version"`
}

type Plan struct {
	Count    int    `json:"count"`
	Interval string `json:"interval"`
	Type     string `json:"type"`
}

type Installments struct {
	Plan Plan `json:"plan"`
}

type VisaAndMaster struct {
	BillingAddress  Address `json:"billing_address"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	ShippingAddress Address `json:"shipping_address"`
}

type Wallet struct {
	AmexExpressCheckout string        `json:"amex_express_checkout"`
	ApplePay            string        `json:"apple_pay"`
	DynamicLast4        string        `json:"dynamic_last4"`
	GooglePay           string        `json:"google_pay"`
	Masterpass          VisaAndMaster `json:"masterpass"`
	SamsungPay          string        `json:"samsung_pay"`
	Type                string        `json:"type"`
	VisaCheckout        VisaAndMaster `json:"visa_checkout"`
}

type Receipt struct {
	AccountType                  string `json:"account_type"`
	ApplicationCryptogram        string `json:"application_cryptogram"`
	ApplicationPreferredName     string `json:"application_preferred_name"`
	AuthorizationCode            string `json:"authorization_code"`
	AuthorizationResponseCode    string `json:"authorization_response_code"`
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
	DedicatedFileName            string `json:"dedicated_file_name"`
	TerminalVerificationResults  string `json:"terminal_verification_results"`
	TransactionStatusInformation string `json:"transaction_status_information"`
}

type Store struct {
	Chain string `json:"chain"`
}
