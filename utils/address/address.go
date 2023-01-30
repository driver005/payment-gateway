package address

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/country"
)

// An address.
type Address struct {
	core.Model

	// Address line 1
	Address1 string `json:"address_1" database:"default:null"`

	// Address line 2
	Address2 string `json:"address_2" database:"default:null"`

	// City
	City string `json:"city" database:"default:null"`

	// Company name
	Company string `json:"company" database:"default:null"`

	// A country object. Available if the relation `country` is expanded.
	Country country.Country `json:"country" database:"foreignKey:id;references:country_code"`

	// The 2 character ISO code of the country in lower case
	CountryCode string `json:"country_code" database:"default:null"`

	// First name
	FirstName string `json:"first_name" database:"default:null"`

	// Last name
	LastName string `json:"last_name" database:"default:null"`

	// Phone Number
	Phone string `json:"phone" database:"default:null"`

	// Postal Code
	PostalCode string `json:"postal_code" database:"default:null"`

	// Province
	Province string `json:"province" database:"default:null"`
}
