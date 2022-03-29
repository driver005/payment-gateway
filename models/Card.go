package models

type Card struct {
	ID                 string      `json:"id" db:"id"`
	Object             string      `json:"object" db:"object"`
	AddressCity        interface{} `json:"address_city" db:"address_city"`
	AddressCountry     interface{} `json:"address_country" db:"address_country"`
	AddressLine1       interface{} `json:"address_line1" db:"address_line1"`
	AddressLine1Check  interface{} `json:"address_line1_check" db:"address_line1_check"`
	AddressLine2       interface{} `json:"address_line2" db:"address_line2"`
	AddressState       interface{} `json:"address_state" db:"address_state"`
	AddressZip         interface{} `json:"address_zip" db:"address_zip"`
	AddressZipCheck    interface{} `json:"address_zip_check" db:"address_zip_check"`
	Brand              string      `json:"brand" db:"brand"`
	Country            string      `json:"country" db:"country"`
	Customer           interface{} `json:"customer" db:"customer"`
	CvcCheck           string      `json:"cvc_check" db:"cvc_check"`
	DynamicLast4       interface{} `json:"dynamic_last4" db:"dynamic_last4"`
	ExpMonth           int         `json:"exp_month" db:"exp_month"`
	ExpYear            int         `json:"exp_year" db:"exp_year"`
	Fingerprint        string      `json:"fingerprint" db:"fingerprint"`
	Funding            string      `json:"funding" db:"funding"`
	Last4              string      `json:"last4" db:"last4"`
	Metadata           struct{}    `json:"metadata"`
	Name               interface{} `json:"name" db:"name"`
	TokenizationMethod interface{} `json:"tokenization_method" db:"tokenization_method"`
}