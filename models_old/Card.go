package models

import (
	"time"

	"github.com/driver005/database"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofrs/uuid"
)

type Card struct {
	ID                 uuid.UUID          `json:"id" db:"id" sql:"default:uuid_generate_v4()"`
	Object             string             `json:"object" db:"object"`
	Address            Address            `json:"address" db:"address" gorm:"type:address"`
	Brand              string             `json:"brand" db:"brand"`
	Country            string             `json:"country" db:"country"`
	Customer           uuid.UUID          `json:"customer" db:"customer"`
	CvcCheck           string             `json:"cvc_check" db:"cvc_check"`
	DynamicLast4       string             `json:"dynamic_last4" db:"dynamic_last4"`
	ExpMonth           int                `json:"exp_month" db:"exp_month"`
	ExpYear            int                `json:"exp_year" db:"exp_year"`
	Fingerprint        string             `json:"fingerprint" db:"fingerprint"`
	Funding            string             `json:"funding" db:"funding"`
	Last4              string             `json:"last4" db:"last4"`
	Metadata           uuid.UUID          `json:"metadata" db:"metadata"`
	Name               string             `json:"name" db:"name"`
	TokenizationMethod string             `json:"tokenization_method" db:"tokenization_method"`
	CreatedAt          time.Time          `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt          database.DeletedAt `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (c *Card) BeforeCreate(tx *database.DB) (err error) {
	c.ID, err = uuid.NewV4()
	if err != nil {
		return err
	}
	return nil
}

func (c Card) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ID, validation.Required, is.UUID),
		validation.Field(&c.Object, validation.Required, validation.In("card")),
		validation.Field(&c.Address, validation.When(c.Address.Validate() != nil)),
		validation.Field(&c.Brand, validation.Required, validation.In("American Express", "Diners Club", "Discover", "JCB", "MasterCard", "UnionPay", "Visa", "Unknown")),
		validation.Field(&c.Country, validation.Required, is.CountryCode2),
		validation.Field(&c.Customer, is.UUID),
		validation.Field(&c.CvcCheck, validation.Required, validation.In("pass", "fail", "unavailable", "unchecked")),
		validation.Field(&c.DynamicLast4),
		validation.Field(&c.ExpMonth, validation.Required, validation.Min(1), validation.Max(12)),
		validation.Field(&c.ExpYear, validation.Required, validation.Min(2000)),
		validation.Field(&c.Fingerprint, validation.Required, validation.Length(16, 32)),
		validation.Field(&c.Funding, validation.Required, validation.In("credit", "debit", "prepaid", "unknown")),
		validation.Field(&c.Last4, validation.Required, validation.Length(4, 4)),
		validation.Field(&c.Name),
		validation.Field(&c.TokenizationMethod, validation.Required, validation.In("android_pay", "apple_pay", "masterpass", "visa_checkout")),
	)
}
