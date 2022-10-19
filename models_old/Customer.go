package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Customer struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Address struct {
		City    string `json:"city" db:"city"`
		State   string `json:"state" db:"state"`
		Country string `json:"country" db:"country"`
		Line1   string `json:"line1" db:"line1"`
		Line2   string `json:"line2" db:"line2"`
		Zip     string `json:"zip" db:"zip"`
	} `json:"address"`
	Balance         int    `json:"balance"`
	Currency        string `json:"currency"`
	DefaultSource   string `json:"default_source"`
	Delinquent      bool   `json:"delinquent"`
	Description     string `json:"description"`
	Discount        string `json:"discount"`
	Email           string `json:"email"`
	InvoicePrefix   string `json:"invoice_prefix"`
	InvoiceSettings struct {
		CustomFields struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"custom_fields"`
		DefaultPaymentMethod string `json:"default_payment_method"`
		Footer               string `json:"footer"`
	} `json:"invoice_settings"`
	Livemode            bool     `json:"livemode"`
	Metadata            string   `json:"metadata"`
	Name                string   `json:"name"`
	NextInvoiceSequence int      `json:"next_invoice_sequence"`
	Phone               string   `json:"phone"`
	PreferredLocales    []string `json:"preferred_locales"`
	Shipping            struct {
		Address struct {
			City    string `json:"city" db:"city"`
			State   string `json:"state" db:"state"`
			Country string `json:"country" db:"country"`
			Line1   string `json:"line1" db:"line1"`
			Line2   string `json:"line2" db:"line2"`
			Zip     string `json:"zip" db:"zip"`
		} `json:"address"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"shipping"`
	TaxExempt string `json:"tax_exempt"`
	TestClock string `json:"test_clock"`
}

func (c Customer) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ID, validation.Required, is.UUID),
		validation.Field(&c.Object, validation.Required, validation.In("Customer")),
		validation.Field(&c.Address.City, validation.Required, validation.Length(5, 50)),
		validation.Field(&c.Address.Country, validation.Required, validation.Length(5, 50), is.CountryCode2),
		validation.Field(&c.Address.Line1, validation.Required, validation.Length(5, 50)),
		validation.Field(&c.Address.Line2, validation.Required, validation.Length(5, 50)),
		validation.Field(&c.Address.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&c.Address.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		validation.Field(&c.Balance, validation.Required),
		validation.Field(&c.Currency, validation.Required, is.CurrencyCode),
		validation.Field(&c.DefaultSource, validation.Required, is.UUID),
		validation.Field(&c.Delinquent, validation.Required),
		validation.Field(&c.Description, validation.Required),
		validation.Field(&c.Discount, is.UUID),
		validation.Field(&c.InvoicePrefix, validation.Required),
		validation.Field(&c.InvoiceSettings.CustomFields.Name),
		validation.Field(&c.InvoiceSettings.CustomFields.Value),
		validation.Field(&c.InvoiceSettings.DefaultPaymentMethod),
		validation.Field(&c.InvoiceSettings.Footer, validation.Required),
		validation.Field(&c.Livemode, validation.Required),
		validation.Field(&c.Metadata, validation.Required, is.UUID),
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.NextInvoiceSequence, validation.Required),
		validation.Field(&c.Phone, validation.Required),
		validation.Field(&c.PreferredLocales, validation.Each(is.CountryCode2)),
		validation.Field(&c.Shipping.Address.City, validation.Length(5, 50)),
		validation.Field(&c.Shipping.Address.Country, validation.Length(5, 50), is.CountryCode2),
		validation.Field(&c.Shipping.Address.Line1, validation.Length(5, 50)),
		validation.Field(&c.Shipping.Address.Line2, validation.Length(5, 50)),
		validation.Field(&c.Shipping.Address.State, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&c.Shipping.Address.Zip, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		validation.Field(&c.Shipping.Name),
		validation.Field(&c.Shipping.Phone),
		validation.Field(&c.TaxExempt, validation.Required, validation.In("none", "exempt", "reverse")),
		validation.Field(&c.TestClock, is.UUID),
	)
}
