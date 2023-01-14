package address

import (
	"regexp"

	"github.com/driver005/database"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
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
	Country *models.Country `json:"country" database:"foreignKey:id;references:country_code"`

	// The 2 character ISO code of the country in lower case
	CountryCode string `json:"country_code" database:"default:null"`

	// Available if the relation `customer` is expanded.
	Customer *models.Customer `json:"customer" database:"foreignKey:id;references:customer_id"`

	// ID of the customer this address belongs to
	CustomerId uuid.NullUUID `json:"customer_id" database:"default:null"`

	// First name
	FirstName string `json:"first_name" database:"default:null"`

	// Last name
	LastName string `json:"last_name" database:"default:null"`

	// An optional key-value map with additional details
	Metadata core.JSONB `json:"metadata" database:"default:null"`

	// Phone Number
	Phone string `json:"phone" database:"default:null"`

	// Postal Code
	PostalCode string `json:"postal_code" database:"default:null"`

	// Province
	Province string `json:"province" database:"default:null"`
}

func (a *Address) BeforeCreate(tx *database.DB) (err error) {
	a.Id, err = uuid.NewUUID()
	if err != nil {
		return err
	}
	return nil
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Id, validation.Required, is.UUID),
		validation.Field(&a.Country, validation.Length(1, 50)),
		validation.Field(&a.Address1, validation.Length(1, 50)),
		validation.Field(&a.Address2, validation.Length(1, 50)),
		validation.Field(&a.City, validation.Length(1, 50)),
		validation.Field(&a.Company, validation.Length(1, 50)),
		validation.Field(&a.Country, validation.Length(1, 50)),
		validation.Field(&a.CountryCode, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&a.FirstName),
		validation.Field(&a.LastName),
		validation.Field(&a.Metadata),
		validation.Field(&a.Phone),
		validation.Field(&a.PostalCode, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		validation.Field(&a.Province, validation.Length(1, 50)),
		// validation.Field(&a.Address1Check, validation.In("pass", "fail", "unavailable", "unchecked")),
		// validation.Field(&a.ZipCheck, validation.In("pass", "fail", "unavailable", "unchecked")),
	)
}

func (a Address) Interface() interface{} {
	return Address(a)
}

// func (Address) DBDataType() string {
// 	return "addressType"
// }

// func (a Address) DBValue(ctx context.Context, db *database.DB) types.Expr {
// 	return types.Expr{
// 		SQL:  "ROW(?, ?, ?, ?, ?, ?, ?, ?)::addressType",
// 		Vars: []interface{}{a.City, a.Country, a.Line1, a.Line1Check, a.Line2, a.State, a.Zip, a.ZipCheck},
// 	}
// }

// func (a *Address) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
// 	}
// 	s := strings.Split(strings.Trim(string(bytes), "()"), ",")
// 	for i := 0; i < reflect.TypeOf(Address{}).NumField(); i++ {
// 		helper.SetField(&a, reflect.TypeOf(Address{}).Field(i).Name, s[i])
// 	}

// 	return nil
// }
