package address

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/country"
)

// An address.
type Address struct {
	core.Model

	// Address line 1
	Address1 string `json:"address_1" database:"default:nul"`

	// Address line 2
	Address2 string `json:"address_2" database:"default:nul"`

	// City
	City string `json:"city" database:"default:nul"`

	// Company name
	Company string `json:"company" database:"default:nul"`

	// A country object. Available if the relation `country` is expanded.
	Country country.Country `json:"country" database:"foreignKey:id;references:country_code"`

	// The 2 character ISO code of the country in lower case
	CountryCode string `json:"country_code" database:"default:nul"`

	// First name
	FirstName string `json:"first_name" database:"default:nul"`

	// Last name
	LastName string `json:"last_name" database:"default:nul"`

	// Phone Number
	Phone string `json:"phone" database:"default:nul"`

	// Postal Code
	PostalCode string `json:"postal_code" database:"default:nul"`

	// Province
	Province string `json:"province" database:"default:nul"`
}
