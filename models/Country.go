package models

import (
	"regexp"

	"github.com/driver005/database"
	"github.com/driver005/gateway/core"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

// Country details
type Country struct {
	core.Model

	// The 2 character ISO code of the country in lower case
	Iso2 string `json:"iso_2"`

	// The 2 character ISO code of the country in lower case
	Iso3 string `json:"iso_3"`

	// The numerical ISO code for the country.
	NumCode string `json:"num_code"`

	// The normalized country name in upper case.
	Name string `json:"name"`

	// The country name appropriate for display.
	DisplayName string `json:"display_name"`

	// The region ID this country is associated with.
	RegionId uuid.NullUUID `json:"region_id" database:"default:null"`

	// A region object. Available if the relation `region` is expanded.
	Region Region `json:"region" database:"default:null"`
}

func (c *Country) BeforeCreate(tx *database.DB) (err error) {
	c.Id, err = uuid.NewUUID()
	if err != nil {
		return err
	}
	return nil
}

func (c Country) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Id, validation.Required, is.UUID),
		validation.Field(&c.DisplayName),
		validation.Field(&c.Name, validation.Match(regexp.MustCompile("[A-Z]+"))),
		validation.Field(&c.Iso2, validation.Match(regexp.MustCompile("^[a-z]{2}$"))),
		validation.Field(&c.Iso3, validation.Match(regexp.MustCompile("^[a-z]{3}$"))),
		validation.Field(&c.NumCode),
		// validation.Field(&c.Address1Check, validation.In("pass", "fail", "unavailable", "unchecked")),
		// validation.Field(&c.ZipCheck, validation.In("pass", "fail", "unavailable", "unchecked")),
	)
}
