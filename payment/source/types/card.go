package types

// SourceTypeCard struct for SourceTypeCard
type SourceTypeCard struct {
	AddressLine1Check string `json:"address_line1_check,omitempty"`
	AddressZipCheck string `json:"address_zip_check,omitempty"`
	Brand string `json:"brand,omitempty"`
	Country string `json:"country,omitempty"`
	CvcCheck string `json:"cvc_check,omitempty"`
	DynamicLast4 string `json:"dynamic_last4,omitempty"`
	ExpMonth int `json:"exp_month,omitempty"`
	ExpYear int `json:"exp_year,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
	Funding string `json:"funding,omitempty"`
	Last4 string `json:"last4,omitempty"`
	Name string `json:"name,omitempty"`
	ThreeDSecure *string `json:"three_d_secure,omitempty"`
	TokenizationMethod string `json:"tokenization_method,omitempty"`
}