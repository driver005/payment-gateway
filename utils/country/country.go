package country

import (
	"github.com/driver005/gateway/core"
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
}
