package models

import (
	"regexp"

	"github.com/driver005/database"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofrs/uuid"
)

// Currency
type Currency struct {
	// The country's ID
	Id uuid.UUID `json:"id,omitempty"`

	// The 3 character ISO code for the currency.
	Code string `json:"code"`

	// [EXPERIMENTAL] Does the currency prices include tax
	IncludesTax *bool `json:"includes_tax,omitempty"`

	// The written name of the currency
	Name string `json:"name"`

	// The symbol used to indicate the currency.
	Symbol string `json:"symbol"`

	// The native symbol used to indicate the currency.
	SymbolNative string `json:"symbol_native"`
}

func (c *Currency) BeforeCreate(tx *database.DB) (err error) {
	c.Id, err = uuid.NewV4()
	if err != nil {
		return err
	}
	return nil
}

func (c Currency) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Id, validation.Required, is.UUID),
		validation.Field(&c.Code, validation.Match(regexp.MustCompile("^[a-z]{3}$"))),
		validation.Field(&c.IncludesTax),
		validation.Field(&c.Name),
		validation.Field(&c.Symbol),
		validation.Field(&c.SymbolNative),
		// validation.Field(&c.Address1Check, validation.In("pass", "fail", "unavailable", "unchecked")),
		// validation.Field(&c.ZipCheck, validation.In("pass", "fail", "unavailable", "unchecked")),
	)
}
