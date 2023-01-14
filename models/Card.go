package models

import (
	"github.com/driver005/database"
	"github.com/driver005/gateway/core"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Card struct {
	core.Model

	Address            *Address  `json:"address"`
	Brand              string    `json:"brand"`
	Country            string    `json:"country"`
	Customer           *Customer `json:"customer"`
	CvcCheck           string    `json:"cvc_check"`
	DynamicLast4       string    `json:"dynamic_last4"`
	ExpMonth           int       `json:"exp_month"`
	ExpYear            int       `json:"exp_year"`
	Fingerprint        string    `json:"fingerprint"`
	Funding            string    `json:"funding"`
	Last4              string    `json:"last4"`
	Name               string    `json:"name"`
	TokenizationMethod string    `json:"tokenization_method"`
}

func (c *Card) BeforeCreate(tx *database.DB) (err error) {
	c.Id, err = uuid.NewUUID()
	if err != nil {
		return err
	}
	return nil
}

func (c *Card) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Id, validation.Required, is.UUID),
		validation.Field(&c.Object, validation.Required, validation.In("card")),
		validation.Field(&c.Address, validation.When(c.Address.Validate() != nil)),
		validation.Field(&c.Brand, validation.Required, validation.In("American Express", "Diners Club", "Discover", "JCB", "MasterCard", "UnionPay", "Visa", "Unknown")),
		validation.Field(&c.Country, validation.Required, is.CountryCode2),
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
