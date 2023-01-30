package types

import "github.com/driver005/gateway/core"

// SourceTypeThreeDSecure struct for SourceTypeThreeDSecure
type SourceTypeThreeDSecure struct {
	core.Model

	AddressLine1Check  string `json:"address_line1_check,omitempty"`
	AddressZipCheck    string `json:"address_zip_check,omitempty"`
	Authenticated      bool   `json:"authenticated,omitempty"`
	Brand              string `json:"brand,omitempty"`
	Card               string `json:"card,omitempty"`
	Country            string `json:"country,omitempty"`
	Customer           string `json:"customer,omitempty"`
	CvcCheck           string `json:"cvc_check,omitempty"`
	DynamicLast4       string `json:"dynamic_last4,omitempty"`
	ExpMonth           int    `json:"exp_month,omitempty"`
	ExpYear            int    `json:"exp_year,omitempty"`
	Fingerprint        string `json:"fingerprint,omitempty"`
	Funding            string `json:"funding,omitempty"`
	Last4              string `json:"last4,omitempty"`
	Name               string `json:"name,omitempty"`
	ThreeDSecure       string `json:"three_d_secure,omitempty"`
	TokenizationMethod string `json:"tokenization_method,omitempty"`
}
